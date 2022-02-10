package class

import (
	"errors"
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

func (this *Class) NewByTag(data contracts.Fields, tag string) interface{} {
	object := reflect.New(this.Type)

	if data != nil {
		for name, field := range this.getFields(tag) {
			if value, exists := data[name]; exists && field.IsExported() {
				object.Elem().FieldByIndex(field.Index).Set(utils.ConvertToValue(field.Type, value))
			}
		}
	}

	return object.Elem().Interface()
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

func (this *Class) New(data contracts.Fields) interface{} {
	return this.NewByTag(data, "json")
}

func (this *Class) getFields(tag string) map[string]reflect.StructField {
	data, exists := this.fields.Load(tag)

	if !exists {
		var fields = map[string]reflect.StructField{}
		for i := 0; i < this.Type.NumField(); i++ {
			field := this.Type.Field(i)
			tags := utils.ParseStructTag(field.Tag)
			if target := tags[tag]; target != nil && len(target) > 0 {
				fields[target[0]] = field
			} else {
				fields[field.Name] = field
			}
		}

		this.fields.Store(tag, fields)
		return fields
	}

	return data.(map[string]reflect.StructField)
}

func (this *Class) ClassName() string {
	return utils.GetTypeKey(this)
}

func (this *Class) GetType() reflect.Type {
	return this.Type
}

func (this *Class) IsSubClass(class interface{}) bool {
	if value, ok := class.(reflect.Type); ok {
		return value.ConvertibleTo(this.Type)
	}

	return reflect.TypeOf(class).ConvertibleTo(this.Type)
}

func (this *Class) Implements(class reflect.Type) bool {
	switch value := class.(type) {
	case *Interface:
		return this.Type.Implements(value.Type)
	case *Class:
		return this.Type.Implements(value.Type)
	}

	return this.Type.Implements(class)
}
