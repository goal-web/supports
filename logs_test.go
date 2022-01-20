package supports_test

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"sync"
	"testing"
)

type TestException struct {
	error
	fields contracts.Fields
}

func (this TestException) Error() string {
	return this.error.Error()
}

func (this TestException) Fields() contracts.Fields {
	return this.fields
}

func TestLogger(t *testing.T) {
	logs.WithError(errors.New("报错了")).Info("info数据")

	logs.WithFields(contracts.Fields{"id": "1"}).Warn("info数据")

	logs.WithException(TestException{
		error:  errors.New("报错啦"),
		fields: contracts.Fields{"id": 1, "name": "qbhy"},
	}).Info("info数据")
}

func TestWithField(t *testing.T) {
	wg := sync.WaitGroup{}
	logs.WithError(errors.New("报错了")).WithField("field1", "1").Info("info数据")

	wg.Add(1)
	go (func() interface{} {
		logs.WithError(errors.New("协程里面报错了")).WithField("field1", "1").Info("info数据")
		wg.Done()
		return nil
	})()

	wg.Wait()
}
