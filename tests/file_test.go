package tests

import (
	"github.com/goal-web/supports/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExists(t *testing.T) {
	assert.True(t, utils.ExistsPath("go.mod"))
}
