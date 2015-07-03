package main

import (
	"github.com/wulonghui/cli"
	_ "github.com/wulonghui/cli/example/commands"
)

func main() {
	cli := cli.New("example", "1.0.0")
	cli.Usage = "This is a example"
	cli.Run()
}
