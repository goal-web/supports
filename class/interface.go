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

func (i *Interface) GetType() reflect.Type {
	return i.Type
}

// Define 创建一个接口
func Define(arg any) contracts.Interface {
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

func (i *Interface) ClassName() string {
	return utils.GetTypeKey(i)
}

func (i *Interface) IsSubClass(class any) bool {
	if value, ok := class.(reflect.Type); ok {
		return value.ConvertibleTo(i.Type)
	}

	return reflect.TypeOf(class).ConvertibleTo(i.Type)
}

func (i *Interface) Implements(class reflect.Type) bool {
	switch value := class.(type) {
	case *Interface:
		return i.Type.Implements(value.Type)
	}

	return i.Type.Implements(class)
}
