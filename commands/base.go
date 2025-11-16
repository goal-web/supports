package commands

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Command struct {
	contracts.CommandArguments
	Signature   string
	Description string
	Name        string
	Help        string
	args        []contracts.CommandArg
}

func Base(signature, description string) *Command {
	name, args := ParseSignature(signature)
	return &Command{
		Signature:   signature,
		Description: description,
		Name:        name,
		Help:        Args(args).Help(),
		args:        args,
	}
}

func (cmd *Command) InjectArguments(defined []contracts.CommandArg, arguments contracts.CommandArguments) error {
	cmd.args = defined
	cmd.CommandArguments = arguments
	argIndex := 0
	for _, arg := range cmd.args {
		switch arg.GetType() {
		case contracts.CommandRequiredArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				if cmd.Exists(arg.GetName()) {
					arguments.SetOption(arg.GetName(), arguments.Get(arg.GetName()))
				} else {
					return fmt.Errorf("missing required parameter：%s - %s", arg.GetName(), arg.GetDescription())
				}
			} else {
				arguments.SetOption(arg.GetName(), argValue)
			}
			argIndex++
		case contracts.CommandOptionalArg:
			argValue := arguments.Optional(arg.GetName(), arg.GetDefault())
			if argValue == "" {
				arguments.SetOption(arg.GetName(), arg.GetDefault())
			} else {
				arguments.SetOption(arg.GetName(), argValue)
			}
			argIndex++
		case contracts.CommandOption:
			if !arguments.Exists(arg.GetName()) && arg.GetDefault() != nil {
				arguments.SetOption(arg.GetName(), arg.GetDefault())
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
func (cmd *Command) GetArgs() []contracts.CommandArg {
	return cmd.args
}
