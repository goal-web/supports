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

func (cmd *Command) InjectArguments(arguments contracts.CommandArguments) error {
	argIndex := 0
	for _, arg := range cmd.args {
		switch arg.Type {
		case RequiredArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				if cmd.Exists(arg.Name) {
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

	cmd.CommandArguments = arguments
	return nil
}

func (cmd *Command) GetSignature() string {
	return cmd.Signature
}
func (cmd *Command) GetDescription() string {
	return cmd.Description
}
func (cmd *Command) GetName() string {
	return cmd.Name
}
func (cmd *Command) GetHelp() string {
	return cmd.Help
}
