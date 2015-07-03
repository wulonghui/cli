package terminal

import (
	"os"
	"time"

	"github.com/codegangsta/cli"
)

var ui UI

func init() {
	ui = NewUI(os.Stdout, NewTeePrinter())
}

func PrintPaginator(rows []string, err error) {
	ui.PrintPaginator(rows, err)
}

func Say(message string, args ...interface{}) {
	ui.Say(message, args...)
}

func PrintCapturingNoOutput(message string, args ...interface{}) {
	ui.PrintCapturingNoOutput(message, args...)
}

func Debug(message string, args ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		ui.Debug(message, args...)
	}
}

func Warn(message string, args ...interface{}) {
	ui.Warn(message, args...)
}

func Ask(prompt string, args ...interface{}) (answer string) {
	return ui.Ask(prompt, args...)
}

func AskForPassword(prompt string, args ...interface{}) (answer string) {
	return ui.AskForPassword(prompt, args...)
}

func Confirm(message string, args ...interface{}) bool {
	return ui.Confirm(message, args...)
}

func ConfirmDelete(modelType, modelName string) bool {
	return ui.ConfirmDelete(modelType, modelName)
}

func ConfirmDeleteWithAssociations(modelType, modelName string) bool {
	return ui.ConfirmDeleteWithAssociations(modelType, modelName)
}

func Ok() {
	ui.Ok()
}

func Failed(message string, args ...interface{}) {
	ui.Failed(message, args...)
}

func FailWithUsage(context *cli.Context) {
	ui.FailWithUsage(context)
}

func PanicQuietly() {
	ui.PanicQuietly()
}

func LoadingIndication() {
	ui.LoadingIndication()
}

func Wait(duration time.Duration) {
	ui.Wait(duration)
}

func PrintTable(headers []string) Table {
	return ui.Table(headers)
}
