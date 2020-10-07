package main

import (
	"fmt"
	"os"
	"time"
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

type line struct {
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
	lines := []*line{}
	messages := []string{
		"Testing 1",
		"Test 2",
		"Testing 3",
	}

	for _, message := range messages {
		lines = append(lines, &line{message: message, status: statusPending})
	}

	ticker := time.NewTicker(200 * time.Millisecond)
	// First print
	for i := range lines {
		fmt.Fprintf(os.Stdout, "%s %s: %s\n", nextTickerChar(), lines[i].message, lines[i].status)
	}
	go func(lines []*line) {
		for {
			<-ticker.C
			for i := range lines {
				fmt.Fprint(os.Stdout, UpLine)
				fmt.Fprintf(os.Stdout, CursorForward(len(lines[i].message)+len(lines[i].status)+padding))
				fmt.Fprintf(os.Stdout, DeleteLine)
			}
			fmt.Fprint(os.Stdout, CursorStartLine)
			tc := nextTickerChar()
			for i := range lines {
				fmt.Fprintf(os.Stdout, "%s %s: %s\n", tc, lines[i].message, lines[i].status)
			}
		}
	}(lines)
	timeout := time.After(10 * time.Second)
	<-timeout
	for i := range lines {
		lines[i].status = statusSuccess
	}
	timeout = time.After(10 * time.Second)
	<-timeout
	for i := range lines {
		lines[i].status = statusFailure
	}
	timeout = time.After(10 * time.Second)
	<-timeout
}
