package main

import (
	"fmt"
	"time"

	"github.com/while1malloc0/multispinner"
)

const (
	UpLine          = "\x1b[1F"
	CursorStartLine = "\r"
	DeleteLine      = "\x1b[1K"
)

const (
	statusPending = "PENDING"
	statusSuccess = "SUCCESS"
	statusFailure = "FAILED"

	padding = 100
)

func CursorForward(n int) string {
	return fmt.Sprintf("\x1b[%dC", n)
}

type row struct {
	status  string
	message string
}

var tickerChars = []string{"\\", "|", "/", "-"}
var currentTickerChar = 0

func nextTickerChar() string {
	currentTickerChar++
	currentTickerChar %= len(tickerChars)
	return tickerChars[currentTickerChar]
}

func main() {
	messages := []string{
		"Testing 1",
		"Test 2",
		"Testing 3",
	}
	spinner := multispinner.NewSpinner()
	rows := []*multispinner.Row{}

	for _, message := range messages {
		row := spinner.AddRow(message)
		rows = append(rows, row)
	}

	spinner.Start()
	timeout := time.After(3 * time.Second)
	<-timeout
	for i := range rows {
		rows[i].Message = "Success message"
	}
	spinner.AddRow("Another row!")
	timeout = time.After(3 * time.Second)
	<-timeout
	for i := range rows {
		rows[i].Message = "Failure message"
	}
	spinner.AddRow("And another row!")
	timeout = time.After(3 * time.Second)
	<-timeout
	spinner.Stop()
}
