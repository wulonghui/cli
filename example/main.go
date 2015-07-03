package main

import (
	"github.com/wulonghui/cli"
	_ "github.com/wulonghui/cli/example/commands"
	. "github.com/wulonghui/cli/i18n"
)

func main() {
	cli := cli.New("example", "1.0.0")
	cli.Usage = T("This is a example")
	cli.Run()
}
