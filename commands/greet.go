package commands

import (
	"fmt"

	"github.com/wulonghui/cli/command"
	. "github.com/wulonghui/cli/i18n"

	"github.com/codegangsta/cli"
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
	name := "someone"
    
	if len(ctx.Args()) > 0 {
		name = ctx.Args()[0]
	}
    
	if ctx.String("lang") == "Spanish" {
		fmt.Println("Hola", name)
	} else {
		fmt.Println("Hello", name)
	}
    
	return nil
}
