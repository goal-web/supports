package supports

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/spf13/cast"
	"strings"
	"time"
)

type InstanceGetter func(key string) any

type BaseFields struct { // 具体方法
	contracts.FieldsProvider // 抽象方法，继承 interface

	Getter InstanceGetter // 如果有设置 getter ，优先使用 getter
}

func (this *BaseFields) Optional(key string, defaultValue any) any {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalTime(key string, defaultValue time.Time) time.Time {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalDuration(key string, defaultValue time.Duration) time.Duration {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalInt8(key string, defaultValue int8) int8 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalInt16(key string, defaultValue int16) int16 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalInt32(key string, defaultValue int32) int32 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalInt64(key string, defaultValue int64) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalUint(key string, defaultValue uint) uint {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalUint8(key string, defaultValue uint8) uint8 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalUint16(key string, defaultValue uint16) uint16 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalUint32(key string, defaultValue uint32) uint32 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalUint64(key string, defaultValue uint64) uint64 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalFloat64(key string, defaultValue float64) float64 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalFloat32(key string, defaultValue float32) float32 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) OptionalBool(key string, defaultValue bool) bool {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetTime(key string) time.Time {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetDuration(key string) time.Duration {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetInt8(key string) int8 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetInt16(key string) int16 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetInt32(key string) int32 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetUint(key string) uint {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetUint8(key string) uint8 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetUint16(key string) uint16 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetUint32(key string) uint32 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetUint64(key string) uint64 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) GetFloat32(key string) float32 {
	//TODO implement me
	panic("implement me")
}

func (this *BaseFields) Get(key string) any {

}

func (this *BaseFields) get(key string) any {
	if this.Getter != nil {
		if value := this.Getter(key); value != nil && value != "" {
			return value
		}
	}
	return this.Fields()[key]
}

func (this *BaseFields) Only(keys ...string) contracts.Fields {
	var fields = make(contracts.Fields)

	for _, key := range keys {
		if value := this.get(key); value != nil {
			fields[key] = value
		}
	}

	return fields
}

func (this *BaseFields) ExceptFields(keys ...string) contracts.Fields {
	var (
		results = make(contracts.Fields)
		keysMap = utils.MakeKeysMap(keys...)
	)

	for key, value := range this.Fields() {
		if _, exists := keysMap[key]; !exists {
			results[key] = value
		}
	}

	return results
}

func (this *BaseFields) OnlyExists(keys ...string) contracts.Fields {
	var fields = make(contracts.Fields)

	for _, key := range keys {
		if value := this.get(key); value != nil {
			fields[key] = value
		}
	}

	return fields
}

func (this *BaseFields) OptionalString(key string, defaultValue string) string {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToString(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) OptionalInt(key string, defaultValue int) int {
	if value := this.get(key); value != nil && value != "" {
		if target, err := cast.ToIntE(value); err == nil {
			return target
		}
		return defaultValue
	}
	return defaultValue
}

func (this *BaseFields) OptionalFields(key string, defaultValue contracts.Fields) contracts.Fields {
	if value := this.get(key); value != nil && value != "" {
		if fields, err := utils.ConvertToFields(value); err == nil {
			return fields
		}
	}
	fields := contracts.Fields{}
	for fieldKey, value := range this.Fields() {
		if strings.HasPrefix(fieldKey, key+".") {
			fields[strings.ReplaceAll(fieldKey, key+".", "")] = value
		}
	}
	if len(fields) > 0 {
		return fields
	}
	return defaultValue
}

func (this *BaseFields) GetString(key string) string {
	return this.OptionalString(key, "")
}

func (this *BaseFields) GetInt64(key string) int64 {
	return this.OptionalInt64(key, 0)
}

func (this *BaseFields) GetInt(key string) int {
	return this.OptionalInt(key, 0)
}

func (this *BaseFields) GetFloat64(key string) float64 {
	return this.OptionalFloat64(key, 0)
}

func (this *BaseFields) GetFloat(key string) float32 {
	return this.OptionalFloat32(key, 0)
}

func (this *BaseFields) GetBool(key string) bool {
	return this.OptionalBool(key, false)
}

func (this *BaseFields) GetFields(key string) contracts.Fields {
	return this.OptionalFields(key, contracts.Fields{})
}
