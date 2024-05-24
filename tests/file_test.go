package tests

import (
	"fmt"
	"github.com/goal-web/supports/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	assert.True(t, utils.ExistsPath("file_test.go"))
}
