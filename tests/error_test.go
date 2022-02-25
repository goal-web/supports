package tests

import (
	"fmt"
	"github.com/goal-web/supports/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoPanic(t *testing.T) {
	var err = utils.NoPanic(func() {
		panic("报错")
	})
	fmt.Println(err)
	assert.Error(t, err)
}
