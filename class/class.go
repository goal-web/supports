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

type Class[T any] struct {
	reflect.Type

	hasPtr bool

	// map[tag名]map[字段名]字段类型
	tagFields sync.Map
	fields    map[string]reflect.StructField
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
		default:
			valueBytes, _ = json.Marshal(v)
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

func (class *Class[T]) NewByTag(data contracts.Fields, tag string) T {
	object := reflect.New(class.Type)
	if !class.hasPtr {
		object = object.Elem()
	}
	assignmentObject := object
	if class.hasPtr {
		assignmentObject = object.Elem()
	}

	if data != nil {
		jsonFields := class.getFields("json")
		targetFields := class.getFields(tag)
		for key, value := range data {
			field, ok := targetFields[key]
			if !ok {
				field, ok = jsonFields[key]
				if !ok {
					field, ok = class.fields[key]
				}
			}
			if ok && field.IsExported() && value != nil {
				jsonValue, isPtr := tryToParseByJson(field.Type, value)
				if jsonValue != nil {
					if isPtr {
						assignmentObject.FieldByIndex(field.Index).Set(reflect.ValueOf(jsonValue))
					} else {
						assignmentObject.FieldByIndex(field.Index).Set(reflect.ValueOf(jsonValue).Elem())
					}
					continue
				} else {
					assignmentObject.FieldByIndex(field.Index).Set(utils.ToValue(field.Type, value))
				}
			}
		}
	}

	return object.Interface().(T)
}

// Make 创建一个类
func Make[T any](args ...T) contracts.Class[T] {
	var arg T
	if len(args) > 0 {
		arg = args[0]
	}
	argType := reflect.TypeOf(arg)
	if argType == nil {
		return nil
	}
	hasPtr := argType.Kind() == reflect.Ptr
	if hasPtr {
		argType = argType.Elem()
	}
	class := &Class[T]{Type: argType, hasPtr: hasPtr}
	if argType.Kind() != reflect.Struct {
		panic(TypeException{Err: errors.New("只支持 struct 类型")})
	}
	class.initFields()
	return class
}

// Any 创建一个类
func Any(arg any) contracts.Class[any] {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Class[any]{Type: argType}
	if argType.Kind() != reflect.Struct {
		panic(TypeException{Err: errors.New("只支持 struct 类型")})
	}
	class.initFields()
	return class
}

func (class *Class[T]) New(data contracts.Fields) T {
	return class.NewByTag(data, "json")
}

func (class *Class[T]) getFields(tag string) map[string]reflect.StructField {
	data, exists := class.tagFields.Load(tag)

	if !exists {
		var fields = map[string]reflect.StructField{}
		for i := 0; i < class.Type.NumField(); i++ {
			field := class.Type.Field(i)
			tags := utils.ParseStructTag(field.Tag)
			if target := tags[tag]; len(target) > 0 {
				fields[target[0]] = field
			} else {
				fields[field.Name] = field
			}
		}

		class.tagFields.Store(tag, fields)
		return fields
	}

	return data.(map[string]reflect.StructField)
}

func (class *Class[T]) initFields() {
	var fields = map[string]reflect.StructField{}
	for i := 0; i < class.Type.NumField(); i++ {
		field := class.Type.Field(i)
		fields[field.Name] = field
	}
	class.fields = fields
}

func (class *Class[T]) ClassName() string {
	return utils.GetTypeKey(class)
}

func (class *Class[T]) GetType() reflect.Type {
	return class.Type
}

func (class *Class[T]) IsSubClass(subclass any) bool {
	if value, ok := subclass.(reflect.Type); ok {
		return value.ConvertibleTo(class.Type)
	}

	return reflect.TypeOf(subclass).ConvertibleTo(class.Type)
}

func (class *Class[T]) Implements(classType reflect.Type) bool {
	switch value := classType.(type) {
	case *Interface:
		return class.Type.Implements(value.Type)
	case *Class[T]:
		return class.Type.Implements(value.Type)
	}

	return class.Type.Implements(classType)
}
