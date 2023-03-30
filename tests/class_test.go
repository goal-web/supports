package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id              int `json:"id" db:"identify"`
	name            string
	Setting         Setting             `json:"setting"`
	SettingStar     *Setting            `json:"setting_star"`
	Settings        []Setting           `json:"settings"`
	MapSettings     map[string]Setting  `json:"map_settings"`
	MapSettingsStar map[string]*Setting `json:"map_settings_star"`
	SettingsStar    []*Setting          `json:"settings_star"`
}

type Setting struct {
	Option string `json:"option"`
}

var UserClass1 = class.Make(new(User))

func TestDefine(t *testing.T) {
	assert.True(t, class.Container.IsSubClass(class.Application))
	assert.True(t, class.Application.Implements(class.Container))
}

func Convert[T any](a any) T {
	return a.(T)
}

func TestConvert(t *testing.T) {
	fmt.Println(Convert[string]("a"))
}

func TestNewByTag(t *testing.T) {
	user := UserClass1.NewByTag(contracts.Fields{
		"identify":          2,
		"name":              "goal", // 为导出字段不支持解析
		"setting":           `{"option": "setting"}`,
		"setting_star":      `{"option": "setting_star"}`,
		"map_settings":      `{"first": {"option": "map_settings"}}`,
		"map_settings_star": `{"first": {"option": "map_settings_star"}}`,
		"settings":          `[{"option": "settings"}]`,
		"settings_star":     `[{"option": "settings_star"}]`,
	}, "db").(User)

	fmt.Println("user.Id", user)
	assert.True(t, user.Id == 2 && user.name == "")
	assert.True(t, user.Setting.Option == "setting")
	assert.True(t, user.SettingStar.Option == "setting_star")
	assert.True(t, len(user.MapSettings) == 1 && user.MapSettings["first"].Option == "map_settings")
	assert.True(t, len(user.MapSettingsStar) == 1 && user.MapSettingsStar["first"].Option == "map_settings_star")
	assert.True(t, len(user.Settings) == 1 && user.Settings[0].Option == "settings")
	assert.True(t, len(user.SettingsStar) == 1 && user.SettingsStar[0].Option == "settings_star")

	user = UserClass1.NewByTag(contracts.Fields{
		"id": 1, // 没有 db 字段没定义，默认就用 json 字段
	}, "db").(User)
	assert.True(t, user.Id == 1)

	var user1 = UserClass1.NewByTag(contracts.Fields{
		"id": []byte("1"), // 没有 db 字段没定义，默认就用 json 字段
	}, "db").(User)
	assert.True(t, user1.Id == 1)
	fmt.Println(user1)
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/supports/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkNewByTag
BenchmarkNewByTag-4   	 2916794	       373.4 ns/op
*/
func BenchmarkNewByTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserClass1.NewByTag(contracts.Fields{
			"identify": 2,
		}, "db")
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/supports/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkComplexNewByTag
BenchmarkComplexNewByTag-4   	 2521668	       491.1 ns/op
*/
func BenchmarkComplexNewByTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := UserClass1.NewByTag(contracts.Fields{
			"identify": i,
		}, "db").(User)
		user.Id++
	}
}
