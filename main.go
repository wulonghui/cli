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

func appHelpTemplate() string {
	return `NAME:
   {{.Name}} - {{.Usage}}
   
USAGE:
   {{.Name}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} [arguments...]
   {{if .Version}}
VERSION:
   {{.Version}}
   {{end}}{{if len .Authors}}
AUTHOR(S): 
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Description}}
   {{end}}{{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}
`
}

func commandHelpTemplate() string {
	return fmt.Sprintf(`NAME:
   {{.Name}} - {{.Description}}
   
USAGE:
   %s {{.Usage}}
   {{if .Flags}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`, NAME)
}

func newApp() (app *cli.App) {
	cli.AppHelpTemplate = appHelpTemplate()
	cli.CommandHelpTemplate = commandHelpTemplate()

	helpCommand := cli.Command{
		Name:        "help",
		Aliases:     []string{"h"},
		Usage:       T("help [COMMAND]"),
		Description: "Shows a list of commands or help for one command",
		Action: func(ctx *cli.Context) {
			args := ctx.Args()
			if args.Present() {
				cli.ShowCommandHelp(ctx, args.First())
			} else {
				cli.ShowAppHelp(ctx)
			}
		},
	}

	app = cli.NewApp()
	app.Usage = usage()
	app.Name = NAME
	app.Version = VERSION
	app.Action = helpCommand.Action
	app.CommandNotFound = func(c *cli.Context, command string) {
		ui.Failed(fmt.Sprintf("Unknown command '%s'", command))
	}

	app.Commands = []cli.Command{helpCommand}

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
