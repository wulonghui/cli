package main

import (
	"fmt"
	"os"

	"github.com/wulonghui/cli/command"
	_ "github.com/wulonghui/cli/commands"
	. "github.com/wulonghui/cli/i18n"
	"github.com/wulonghui/cli/terminal"

	"github.com/codegangsta/cli"
)

var ui terminal.UI

func init() {
	ui = terminal.NewUI(os.Stdout, terminal.NewTeePrinter())
}

func runCommand(cmd command.Command, ctx *cli.Context) (err error) {
	cmd, err = cmd.Init(command.CommandInitData{
		UI: ui,
	})

	if err != nil {
		return err
	}

	err = cmd.Execute(ctx)

	return err
}

func usage() string {
	return T("A command line tool")
}

func newApp() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = usage()
	app.Name = NAME
	app.Version = VERSION
	app.CommandNotFound = func(c *cli.Context, command string) {
		ui.Failed(fmt.Sprintf("Unknown command '%s'", command))
	}

	app.Commands = make([]cli.Command, 0)

	for _, cmd := range command.Commands {
		metadata := cmd.MetaData()
		app.Commands = append(app.Commands, cli.Command{
			Name:            metadata.Name,
			ShortName:       metadata.ShortName,
			Description:     metadata.Description,
			Usage:           metadata.Usage,
			Flags:           metadata.Flags,
			SkipFlagParsing: metadata.SkipFlagParsing,
			Action: func(ctx *cli.Context) {
				err := runCommand(cmd, ctx)
				if err != nil {
					panic(err)
				}
			},
		})
	}

	return
}

func main() {
	defer handlePanics()

	newApp().Run(os.Args)
}
