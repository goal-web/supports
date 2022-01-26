package tests

import (
	"github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefine(t *testing.T) {
	assert.True(t, class.Container.IsSubClass(class.Application))
	assert.True(t, class.Application.Implements(class.Container))
}
