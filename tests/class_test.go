package tests

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id int `db:"identify"`
}

var UserClass1 = class.Make(new(User))

func TestDefine(t *testing.T) {
	assert.True(t, class.Container.IsSubClass(class.Application))
	assert.True(t, class.Application.Implements(class.Container))
}

func TestNewByTag(t *testing.T) {
	user := UserClass1.NewByTag(contracts.Fields{
		"id":       1,
		"Id":       1,
		"identify": 2,
	}, "db").(User)

	assert.True(t, user.Id == 2)
}

/**
goos: darwin
goarch: amd64
pkg: github.com/goal-web/supports/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkNewByTag
BenchmarkNewByTag-4   	 2717953	       388.3 ns/op
*/
func BenchmarkNewByTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserClass1.NewByTag(contracts.Fields{
			"identify": 2,
		}, "db")
	}
}

/**
goos: darwin
goarch: amd64
pkg: github.com/goal-web/supports/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkComplexNewByTag
BenchmarkComplexNewByTag-4   	 2704674	       404.3 ns/op
*/
func BenchmarkComplexNewByTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := UserClass1.NewByTag(contracts.Fields{
			"identify": i,
		}, "db").(User)
		user.Id++
	}
}
