package command

import (
	"github.com/codegangsta/cli"
)

type Command interface {
	MetaData() CommandMetadata
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
