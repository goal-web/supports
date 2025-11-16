package commands

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/modood/table"
	"regexp"
	"strings"
)

type Arg struct {
	Name        string
	Type        contracts.CommandArgType
	Default     any
	Description string
}

func (a Arg) GetName() string {
	return a.Name
}

func (a Arg) GetType() contracts.CommandArgType {
	return a.Type
}

func (a Arg) GetDefault() any {
	return a.Default
}

func (a Arg) GetDescription() string {
	return a.Description
}

type Args []contracts.CommandArg

func (args Args) Help() string {
	if len(args) > 0 {
		return table.Table(args)
	}
	return "该命令无参数"
}

func NewArg(name string, argType contracts.CommandArgType, defaultValue any) Arg {
	if names := strings.Split(name, ":"); len(names) > 1 { // 有定义描述
		return Arg{
			Name:        names[0],
			Type:        argType,
			Default:     defaultValue,
			Description: names[1],
		}
	} else {
		return Arg{
			Name:        name,
			Type:        argType,
			Default:     defaultValue,
			Description: "",
		}
	}
}

func ParseSignature(signature string) (string, []contracts.CommandArg) {
	cmd := strings.Split(signature, " ")[0]
	reg, _ := regexp.Compile(" {([^{}]*)}")
	args := make([]contracts.CommandArg, 0)

	for _, arg := range reg.FindAllString(signature, -1) {
		arg = utils.SubString(arg, 2, -1)
		if argArr := strings.Split(arg, "="); len(argArr) > 1 { // {name=goal} / {--name=goal}
			if strings.HasPrefix(argArr[0], "--") {
				args = append(args, NewArg(utils.SubString(argArr[0], 2, 0), contracts.CommandOption, argArr[1]))
			} else {
				args = append(args, NewArg(argArr[0], contracts.CommandOptionalArg, argArr[1]))
			}
		} else if strings.HasSuffix(arg, "?") { // {name?}
			arg = utils.SubString(arg, 0, -1)
			args = append(args, NewArg(arg, contracts.CommandOptionalArg, nil))
		} else if strings.HasPrefix(arg, "--") { // {--name}
			arg = utils.SubString(arg, 2, 0)
			args = append(args, NewArg(arg, contracts.CommandOption, nil))
		} else { // {name}
			args = append(args, NewArg(arg, contracts.CommandRequiredArg, nil))
		}
	}
	return cmd, args
}
