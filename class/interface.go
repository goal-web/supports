package class

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/pkg/errors"
	"reflect"
)

type Interface struct {
	reflect.Type

	// map[字段名]字段类型
	fields map[string]reflect.StructField
}

func (this *Interface) GetType() reflect.Type {
	return this.Type
}

// Define 创建一个接口
func Define(arg interface{}) contracts.Interface {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Interface{Type: argType}
	if argType.Kind() != reflect.Interface {
		panic(TypeException{Err: errors.New("只支持 interface 类型")})
	}
	return class
}

func (this *Interface) ClassName() string {
	return utils.GetTypeKey(this)
}

func (this *Interface) IsSubClass(class interface{}) bool {
	if value, ok := class.(reflect.Type); ok {
		return value.ConvertibleTo(this.Type)
	}

	return reflect.TypeOf(class).ConvertibleTo(this.Type)
}

func (this *Interface) Implements(class reflect.Type) bool {
	switch value := class.(type) {
	case *Interface:
		return this.Type.Implements(value.Type)
	case *Class:
		return false
	}

	return this.Type.Implements(class)
}
