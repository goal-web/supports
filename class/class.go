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

func (class *Class) NewByTag(data contracts.Fields, tag string) interface{} {
	object := reflect.New(class.Type).Elem()

	if data != nil {
		jsonFields := class.getFields("json")
		targetFields := class.getFields(tag)
		for key, value := range data {
			field, ok := jsonFields[key]
			fieldExported := field.IsExported()
			if ok && fieldExported && (field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct)) {
				var valueBytes []byte
				switch v := value.(type) {
				case []byte:
					valueBytes = v
				case string:
					valueBytes = []byte(v)
				case fmt.Stringer:
					valueBytes = []byte(v.String())
				default:
					valueBytes = []byte(utils.ConvertToString(v, ""))
				}
				var fieldValue interface{}
				var isStruct = field.Type.Kind() == reflect.Struct
				if isStruct {
					fieldValue = reflect.New(field.Type).Interface()
				} else {
					fieldValue = reflect.New(field.Type.Elem()).Interface()
				}
				err := json.Unmarshal(valueBytes, &fieldValue)
				if err == nil {
					if isStruct {
						object.FieldByIndex(field.Index).Set(reflect.ValueOf(fieldValue).Elem())
					} else {
						object.FieldByIndex(field.Index).Set(reflect.ValueOf(fieldValue))
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
func Make(arg interface{}) contracts.Class {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Class{Type: argType}
	if argType.Kind() != reflect.Struct {
		panic(TypeException{
			errors.New("只支持 struct 类型"),
			map[string]interface{}{
				"class": class.ClassName(),
			},
		})
	}
	return class
}

func (class *Class) New(data contracts.Fields) interface{} {
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

func (class *Class) IsSubClass(subclass interface{}) bool {
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
