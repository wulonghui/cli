package commands

import (
	"github.com/codegangsta/cli"
	"github.com/wulonghui/cli/command"
	. "github.com/wulonghui/cli/i18n"
	"github.com/wulonghui/cli/terminal"
)

func init() {
	command.Register(Greet{})
}

type Greet struct {
}

func (cmd Greet) MetaData() command.CommandMetadata {
	return command.CommandMetadata{
		Name:        "greet",
		ShortName:   "g",
		Usage:       T("greet [NAME]"),
		Description: T("Say hello"),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "lang, l",
				Value: T("English or Spanish, default is English"),
				Usage: T("Language for the greeting"),
			},
		},
	}
}

func (cmd Greet) Execute(ctx *cli.Context) error {
	terminal.Debug("Greet %v", ctx.Args())

	var name string

	if len(ctx.Args()) == 0 {
		name = terminal.Ask(T("Input name"))
	} else if len(ctx.Args()) == 1 {
		name = ctx.Args()[0]
	} else {
		terminal.FailWithUsage(ctx)
	}

	if ctx.String("lang") == "Spanish" {
		terminal.Say("Hola %s", name)
	} else {
		terminal.Say("Hello %s", name)
	}

	return nil
}
