package command

import (
	"github.com/codegangsta/cli"
	"github.com/wulonghui/cli/terminal"
)

type Command interface {
	MetaData() CommandMetadata
	Init(data CommandInitData) (Command, error)
	Execute(ctx *cli.Context) error
}

type CommandMetadata struct {
	Name            string
	ShortName       string
	Usage           string
	Description     string
	Flags           []cli.Flag
	SkipFlagParsing bool
}

type CommandInitData struct {
	UI terminal.UI
}
