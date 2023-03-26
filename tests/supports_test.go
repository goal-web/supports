package tests

import (
	"fmt"
	"github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type UserClass struct {
	Name string `json:"name"`
}

func TestClass(t *testing.T) {
	class := class.Make(UserClass{})

	userInstance := class.New(map[string]any{
		"name": "goal",
	}).(UserClass)

	fmt.Println(userInstance)

	assert.True(t, userInstance.Name == "goal")
}
