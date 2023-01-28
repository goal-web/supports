package tests_test

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"sync"
	"testing"
)

type TestException = exceptions.Exception

func TestLogger(t *testing.T) {
	logs.WithError(errors.New("报错了")).Info("info数据")

	logs.WithFields(contracts.Fields{"id": "1"}).Warn("info数据")

	logs.WithException(&TestException{Err: errors.New("报错啦")}).Info("info数据")
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
