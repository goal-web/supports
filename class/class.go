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

func (c *Class) NewByTag(data contracts.Fields, tag string) any {
	object := reflect.New(c.Type).Elem()

	if data != nil {
		jsonFields := c.getFields("json")
		targetFields := c.getFields(tag)
		for key, value := range data {
			if field, ok := targetFields[key]; ok && field.IsExported() {
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

func (c *Class) New(data contracts.Fields) any {
	return c.NewByTag(data, "json")
}

func (c *Class) getFields(tag string) map[string]reflect.StructField {
	data, exists := c.fields.Load(tag)

	if !exists {
		var fields = map[string]reflect.StructField{}
		for i := 0; i < c.Type.NumField(); i++ {
			field := c.Type.Field(i)
			tags := utils.ParseStructTag(field.Tag)
			if target := tags[tag]; target != nil && len(target) > 0 {
				fields[target[0]] = field
			} else {
				fields[field.Name] = field
			}
		}

		c.fields.Store(tag, fields)
		return fields
	}

	return data.(map[string]reflect.StructField)
}

func (c *Class) ClassName() string {
	return utils.GetTypeKey(c)
}

func (c *Class) GetType() reflect.Type {
	return c.Type
}

func (c *Class) IsSubClass(class any) bool {
	if value, ok := class.(reflect.Type); ok {
		return value.ConvertibleTo(c.Type)
	}

	return reflect.TypeOf(class).ConvertibleTo(c.Type)
}

func (c *Class) Implements(class reflect.Type) bool {
	switch value := class.(type) {
	case *Interface:
		return c.Type.Implements(value.Type)
	case *Class:
		return c.Type.Implements(value.Type)
	}

	return c.Type.Implements(class)
}
