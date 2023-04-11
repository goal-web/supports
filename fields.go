package supports

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

type InstanceGetter func(key string) any

type BaseFields struct { // 具体方法
	contracts.FieldsProvider // 抽象方法，继承 interface

	Getter InstanceGetter // 如果有设置 getter ，优先使用 getter
}

func (base *BaseFields) Optional(key string, value any) any {
	if result := base.get(key); result != nil {
		return result
	}
	return value
}

func (base *BaseFields) Int16Optional(key string, defaultValue int16) int16 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToInt16(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) Int32Optional(key string, defaultValue int32) int32 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToInt32(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) Int8Optional(key string, defaultValue int8) int8 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToInt8(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) UInt64Optional(key string, defaultValue uint64) uint64 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToUInt64(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) UInt32Optional(key string, defaultValue uint32) uint32 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToUInt32(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) UInt16Optional(key string, defaultValue uint16) uint16 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToUInt16(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) UInt8Optional(key string, defaultValue uint8) uint8 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToUInt8(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) UIntOptional(key string, defaultValue uint) uint {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToUInt(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) get(key string) any {
	if base.Getter != nil {
		if value := base.Getter(key); value != nil && value != "" {
			return value
		}
	}
	return base.Fields()[key]
}

func (base *BaseFields) Only(keys ...string) contracts.Fields {
	var fields = make(contracts.Fields)

	for _, key := range keys {
		if value := base.get(key); value != nil {
			fields[key] = value
		}
	}

	return fields
}

func (base *BaseFields) ExceptFields(keys ...string) contracts.Fields {
	var (
		results = make(contracts.Fields)
		keysMap = utils.MakeKeysMap(keys...)
	)

	for key, value := range base.Fields() {
		if _, exists := keysMap[key]; !exists {
			results[key] = value
		}
	}

	return results
}

func (base *BaseFields) OnlyExists(keys ...string) contracts.Fields {
	var fields = make(contracts.Fields)

	for _, key := range keys {
		if value := base.get(key); value != nil {
			fields[key] = value
		}
	}

	return fields
}

func (base *BaseFields) StringOptional(key string, defaultValue string) string {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToString(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) Int64Optional(key string, defaultValue int64) int64 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToInt64(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) IntOptional(key string, defaultValue int) int {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToInt(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) Float64Optional(key string, defaultValue float64) float64 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToFloat64(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) FloatOptional(key string, defaultValue float32) float32 {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToFloat(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) BoolOptional(key string, defaultValue bool) bool {
	if value := base.get(key); value != nil && value != "" {
		return utils.ToBool(value, defaultValue)
	}
	return defaultValue
}

func (base *BaseFields) GetString(key string) string {
	return base.StringOptional(key, "")
}

func (base *BaseFields) GetInt64(key string) int64 {
	return base.Int64Optional(key, 0)
}

func (base *BaseFields) GetInt(key string) int {
	return base.IntOptional(key, 0)
}

func (base *BaseFields) GetFloat64(key string) float64 {
	return base.Float64Optional(key, 0)
}

func (base *BaseFields) GetFloat(key string) float32 {
	return base.FloatOptional(key, 0)
}

func (base *BaseFields) GetBool(key string) bool {
	return base.BoolOptional(key, false)
}

func (base *BaseFields) Get(key string) any {
	return base.get(key)
}

func (base *BaseFields) GetInt32(key string) int32 {
	return base.Int32Optional(key, 0)
}

func (base *BaseFields) GetInt16(key string) int16 {
	return base.Int16Optional(key, 0)
}

func (base *BaseFields) GetInt8(key string) int8 {
	return base.Int8Optional(key, 0)
}

func (base *BaseFields) GetUInt64(key string) uint64 {
	return base.UInt64Optional(key, 0)
}

func (base *BaseFields) GetUInt32(key string) uint32 {
	return base.UInt32Optional(key, 0)
}

func (base *BaseFields) GetUInt16(key string) uint16 {
	return base.UInt16Optional(key, 0)
}

func (base *BaseFields) GetUInt8(key string) uint8 {
	return base.UInt8Optional(key, 0)
}

func (base *BaseFields) GetUInt(key string) uint {
	return base.UIntOptional(key, 0)
}
