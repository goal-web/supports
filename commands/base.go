package commands

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
)

type Command struct {
	contracts.CommandArguments
	Signature   string
	Description string
	Name        string
	Help        string
	args        []Arg
}

func Base(signature, description string) Command {
	name, args := ParseSignature(signature)
	return Command{
		Signature:   signature,
		Description: description,
		Name:        name,
		Help:        args.Help(),
		args:        args,
	}
}

func (this *Command) InjectArguments(arguments contracts.CommandArguments) error {
	argIndex := 0
	for _, arg := range this.args {
		switch arg.Type {
		case RequiredArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				if this.Exists(arg.Name) {
					arguments.SetOption(arg.Name, arguments.Fields()[arg.Name])
				} else {
					return errors.New(fmt.Sprintf("Missing required parameter：%s - %s", arg.Name, arg.Description))
				}
			} else {
				arguments.SetOption(arg.Name, argValue)
			}
			argIndex++
		case OptionalArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				arguments.SetOption(arg.Name, arg.Default)
			} else {
				arguments.SetOption(arg.Name, argValue)
			}
			argIndex++
		case Option:
			if !arguments.Exists(arg.Name) && arg.Default != nil {
				arguments.SetOption(arg.Name, arg.Default)
			}
		}
	}

	this.CommandArguments = arguments
	return nil
}

func (this *Command) GetSignature() string {
	return this.Signature
}
func (this *Command) GetDescription() string {
	return this.Description
}
func (this *Command) GetName() string {
	return this.Name
}
func (this *Command) GetHelp() string {
	return this.Help
}
