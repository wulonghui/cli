package commands

import (
	"github.com/wulonghui/cli/command"
	. "github.com/wulonghui/cli/i18n"
	"github.com/wulonghui/cli/terminal"

	"github.com/codegangsta/cli"
)

func init() {

	command.Register(Greet{})
}

type Greet struct {
	terminal.UI
}

func (cmd Greet) Init(data command.CommandInitData) (command.Command, error) {
	cmd.UI = data.UI
	return cmd, nil
}

func (cmd Greet) MetaData() command.CommandMetadata {
	return command.CommandMetadata{
		Name:        "greet",
		ShortName:   "g",
		Usage:       T("Say hello"),
		Description: T("Say hello"),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "lang, l",
				Value: "English or Spanish, default is English",
				Usage: "Language for the greeting",
			},
		},
	}
}

func (cmd Greet) Execute(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		cmd.FailWithUsage(ctx)
	}

	name := ctx.Args()[0]

	if ctx.String("lang") == "Spanish" {
		cmd.Say("Hola %s", name)
	} else {
		cmd.Say("Hello %s", name)
	}

	return nil
}
