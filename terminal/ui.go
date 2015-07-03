package terminal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	. "github.com/wulonghui/cli/i18n"

	"github.com/codegangsta/cli"
)

type ColoringFunction func(value string, row int, col int) string

type UI interface {
	PrintPaginator(rows []string, err error)
	Say(message string, args ...interface{})
	PrintCapturingNoOutput(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Ask(prompt string, args ...interface{}) (answer string)
	AskForPassword(prompt string, args ...interface{}) (answer string)
	Confirm(message string, args ...interface{}) bool
	ConfirmDelete(modelType, modelName string) bool
	ConfirmDeleteWithAssociations(modelType, modelName string) bool
	Ok()
	Failed(message string, args ...interface{})
	FailWithUsage(context *cli.Context)
	PanicQuietly()
	LoadingIndication()
	Wait(duration time.Duration)
	Table(headers []string) Table
}

type terminalUI struct {
	stdin   io.Reader
	printer Printer
}

func NewUI(r io.Reader, printer Printer) UI {
	return &terminalUI{
		stdin:   r,
		printer: printer,
	}
}

func (c *terminalUI) PrintPaginator(rows []string, err error) {
	if err != nil {
		c.Failed(err.Error())
		return
	}

	for _, row := range rows {
		c.Say(row)
	}
}

func (c *terminalUI) PrintCapturingNoOutput(message string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Printf("%s", message)
	} else {
		fmt.Printf(message, args...)
	}
}

func (c *terminalUI) Say(message string, args ...interface{}) {
	if len(args) == 0 {
		c.printer.Printf("%s\n", message)
	} else {
		c.printer.Printf(message+"\n", args...)
	}
}

func (c *terminalUI) Debug(message string, args ...interface{}) {
	message = fmt.Sprintf("[Debug] "+message, args...)
	c.Say(DebugColor(message))
	return
}

func (c *terminalUI) Warn(message string, args ...interface{}) {
	message = fmt.Sprintf("[WARN] "+message, args...)
	c.Say(WarningColor(message))
	return
}

func (c *terminalUI) ConfirmDeleteWithAssociations(modelType, modelName string) bool {
	return c.confirmDelete(T("Really delete the {{.ModelType}} {{.ModelName}} and everything associated with it?",
		map[string]interface{}{
			"ModelType": modelType,
			"ModelName": EntityNameColor(modelName),
		}))
}

func (c *terminalUI) ConfirmDelete(modelType, modelName string) bool {
	return c.confirmDelete(T("Really delete the {{.ModelType}} {{.ModelName}}?",
		map[string]interface{}{
			"ModelType": modelType,
			"ModelName": EntityNameColor(modelName),
		}))
}

func (c *terminalUI) confirmDelete(message string) bool {
	result := c.Confirm(message)

	if !result {
		c.Warn(T("Delete cancelled"))
	}

	return result
}

func (c *terminalUI) Confirm(message string, args ...interface{}) bool {
	response := c.Ask(message, args...)
	switch strings.ToLower(response) {
	case "y", "yes", T("yes"):
		return true
	}
	return false
}

func (c *terminalUI) Ask(prompt string, args ...interface{}) (answer string) {
	fmt.Println("")
	fmt.Printf(prompt+PromptColor(">")+" ", args...)

	rd := bufio.NewReader(c.stdin)
	line, err := rd.ReadString('\n')
	if err == nil {
		return strings.TrimSpace(line)
	}
	return ""
}

func (c *terminalUI) Ok() {
	c.Say(SuccessColor(T("OK")))
}

const QuietPanic = "This shouldn't print anything"

func (c *terminalUI) Failed(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)

	if T == nil {
		c.Say(FailureColor("FAILED"))
		c.Say(message)

		c.PanicQuietly()
	} else {
		c.Say(FailureColor(T("FAILED")))
		c.Say(message)

		c.PanicQuietly()
	}
}

func (c *terminalUI) PanicQuietly() {
	panic(QuietPanic)
}

func (c *terminalUI) FailWithUsage(context *cli.Context) {
	c.Say(FailureColor(T("FAILED")))
	c.Say(T("Incorrect Usage.\n"))
	cli.ShowCommandHelp(context, context.Command.Name)
	c.Say("")
	os.Exit(1)
}

func (c *terminalUI) LoadingIndication() {
	c.printer.Print(".")
}

func (c *terminalUI) Wait(duration time.Duration) {
	time.Sleep(duration)
}

func (ui *terminalUI) Table(headers []string) Table {
	return NewTable(ui, headers)
}
