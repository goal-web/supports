package class

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"reflect"
)

type Class struct {
	reflect.Type

	// map[字段名]字段类型
	fields map[string]reflect.StructField
}

type Exception struct {
	error
	fields contracts.Fields
}

func (this Exception) Error() string {
	return this.error.Error()
}

func (this Exception) Fields() contracts.Fields {
	return this.fields
}

// Make 创建一个类
func Make(arg interface{}) *Class {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Class{Type: argType}
	if argType.Kind() != reflect.Struct {
		panic(Exception{
			errors.New("只支持 struct 类型"),
			map[string]interface{}{
				"class": class.ClassName(),
			},
		})
	}
	class.init()
	return class
}

func (this *Class) New(data contracts.Fields) interface{} {
	object := reflect.New(this.Type)

	for name, field := range this.fields {
		if value, exists := data[name]; exists && field.IsExported() {
			object.Elem().FieldByIndex(field.Index).Set(utils.ConvertToValue(field.Type, value))
		}
	}

	return object.Elem().Interface()
}

func (this *Class) ClassName() string {
	return utils.GetTypeKey(this)
}

func (this *Class) init() {
	if this.fields == nil {
		this.fields = map[string]reflect.StructField{}
	}
	for i := 0; i < this.Type.NumField(); i++ {
		field := this.Type.Field(i)
		tags := utils.ParseStructTag(field.Tag)
		if jsonTag := tags["json"]; jsonTag != nil && len(jsonTag) > 0 {
			this.fields[jsonTag[0]] = field
		} else {
			this.fields[field.Name] = field
		}
	}
}
