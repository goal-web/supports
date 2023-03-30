package class

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"reflect"
	"sync"
)

type Class struct {
	reflect.Type

	// map[tag名]map[字段名]字段类型
	fields sync.Map
}

func tryToParseByJson(t reflect.Type, value interface{}) (interface{}, bool) {
	kind := t.Kind()
	var maybeCanJsonify bool
	switch kind {
	case reflect.Struct, reflect.Array, reflect.Map, reflect.Slice:
		maybeCanJsonify = true
	case reflect.Ptr:
		switch t.Elem().Kind() {
		case reflect.Struct, reflect.Array, reflect.Map, reflect.Slice:
			maybeCanJsonify = true
		}
	}
	var isPtr = kind == reflect.Ptr
	if maybeCanJsonify {
		var valueBytes []byte
		switch v := value.(type) {
		case []byte:
			valueBytes = v
		case string:
			valueBytes = []byte(v)
		case fmt.Stringer:
			valueBytes = []byte(v.String())
		}

		if len(valueBytes) > 0 {
			if isPtr {
				fieldValue := reflect.New(t.Elem()).Interface()
				err := json.Unmarshal(valueBytes, &fieldValue)
				if err == nil {
					return fieldValue, isPtr
				}
			} else {
				fieldValue := reflect.New(t).Interface()
				err := json.Unmarshal(valueBytes, &fieldValue)
				if err == nil {
					return fieldValue, isPtr
				}
			}
		}
	}
	return nil, isPtr
}

func (class *Class) NewByTag(data contracts.Fields, tag string) any {
	object := reflect.New(class.Type).Elem()

	if data != nil {
		jsonFields := class.getFields("json")
		targetFields := class.getFields(tag)
		for key, value := range data {
			field, ok := jsonFields[key]
			fieldExported := field.IsExported()
			if ok && fieldExported {
				jsonValue, isPtr := tryToParseByJson(field.Type, value)
				if jsonValue != nil {
					if isPtr {
						object.FieldByIndex(field.Index).Set(reflect.ValueOf(jsonValue))
					} else {
						object.FieldByIndex(field.Index).Set(reflect.ValueOf(jsonValue).Elem())
					}
					continue
				}
			}

			if field, ok = targetFields[key]; ok && field.IsExported() {
				object.FieldByIndex(field.Index).Set(utils.ConvertToValue(field.Type, value))
			} else if field, ok = jsonFields[key]; ok && field.IsExported() {
				object.FieldByIndex(field.Index).Set(utils.ConvertToValue(field.Type, value))
			}
		}
	}

	return object.Interface()
}

// Make 创建一个类
func Make(arg any) contracts.Class {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Class{Type: argType}
	if argType.Kind() != reflect.Struct {
		panic(TypeException{Err: errors.New("只支持 struct 类型")})
	}
	return class
}

func (class *Class) New(data contracts.Fields) any {
	return class.NewByTag(data, "json")
}

func (class *Class) getFields(tag string) map[string]reflect.StructField {
	data, exists := class.fields.Load(tag)

	if !exists {
		var fields = map[string]reflect.StructField{}
		for i := 0; i < class.Type.NumField(); i++ {
			field := class.Type.Field(i)
			tags := utils.ParseStructTag(field.Tag)
			if target := tags[tag]; target != nil && len(target) > 0 {
				fields[target[0]] = field
			} else {
				fields[field.Name] = field
			}
		}

		class.fields.Store(tag, fields)
		return fields
	}

	return data.(map[string]reflect.StructField)
}

func (class *Class) ClassName() string {
	return utils.GetTypeKey(class)
}

func (class *Class) GetType() reflect.Type {
	return class.Type
}

func (class *Class) IsSubClass(subclass any) bool {
	if value, ok := subclass.(reflect.Type); ok {
		return value.ConvertibleTo(class.Type)
	}

	return reflect.TypeOf(subclass).ConvertibleTo(class.Type)
}

func (class *Class) Implements(classType reflect.Type) bool {
	switch value := classType.(type) {
	case *Interface:
		return class.Type.Implements(value.Type)
	case *Class:
		return class.Type.Implements(value.Type)
	}

	return class.Type.Implements(classType)
}
