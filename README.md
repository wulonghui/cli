# cli
This is a command line tool package wrapper around [codegangsta/cli](https://github.com/codegangsta/cli)  to help you easy to create a cli famework, also reference [cloudfoundry/cli](https://github.com/cloudfoundry/cli) to provide some useful tools for cli(.e.g terminal output and interaction, i18n ...).

##Usage
There is a [example](https://github.com/wulonghui/cli/tree/master/example):
```
NAME:
   example - This is a example
   
USAGE:
   example [global options] command [command options] [arguments...]
   
VERSION:
   1.0.0
   
COMMANDS:
   help, h	Shows a list of commands or help for one command
   greet, g	Say hello
   
GLOBAL OPTIONS:
   --debug, -d		Enable debug mode [$DEBUG]
   --version, -v	print the version
```

First new a cli:
```
package main

import (
	"github.com/wulonghui/cli"
	_ "github.com/wulonghui/cli/example/commands"
)

func main(){
    cli := cli.New("example", "1.0.0")
    cli.Usage = "This is a example"
    cli.Run()
}

```
and then add command:
```
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
				Value: "English or Spanish, default is English",
				Usage: "Language for the greeting",
			},
		},
	}
}

func (cmd Greet) Execute(ctx *cli.Context) error {
	terminal.Debug("Greet %v", ctx.Args())
    
    var name string
    
	if len(ctx.Args()) == 0 {
		name = terminal.Ask("Input name")
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

```

