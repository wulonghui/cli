package main

import (
	"fmt"
	"os"

	"github.com/wulonghui/cli/command"
	_ "github.com/wulonghui/cli/commands"
	. "github.com/wulonghui/cli/i18n"

	"github.com/codegangsta/cli"
)

func usage() string {
	return T("A command line tool")
}

func NewApp() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = usage()
	app.Name = NAME
	app.Version = VERSION
	app.CommandNotFound = func(c *cli.Context, command string) {
		panic(fmt.Sprintf("'%s' is not a registered command", command))
	}

	app.Commands = make([]cli.Command, 0)

	for _, c := range command.Commands {
		metadata := c.MetaData()
		app.Commands = append(app.Commands, cli.Command{
			Name:            metadata.Name,
			ShortName:       metadata.ShortName,
			Description:     metadata.Description,
			Usage:           metadata.Usage,
			Flags:           metadata.Flags,
			SkipFlagParsing: metadata.SkipFlagParsing,
			Action: func(ctx *cli.Context) {
				err := c.Execute(ctx)
				if err != nil {
					panic(err)
				}
			},
		})
	}

	return
}

func handlePanics() {
	err := recover()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	defer handlePanics()

	NewApp().Run(os.Args)
}
