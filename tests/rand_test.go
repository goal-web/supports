package tests

import (
	"fmt"
	"github.com/goal-web/supports/utils"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 50; i++ {
		fmt.Println(utils.RandInt(-100, 200))
	}
}

func TestRandIntArray(t *testing.T) {
	for i := 0; i < 50; i++ {
		fmt.Println(utils.RandIntArray(-100, 200, 10))
	}
}
