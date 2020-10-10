package multispinner

import (
	"fmt"
	"os"
	"time"
)

const (
	PendingStatus = iota
	SuccessStatus
	FailedStatus

	UpLine          = "\x1b[1F"
	CursorStartLine = "\r"
	DeleteLine      = "\x1b[1K"

	padding = 100
)

type Status int

type Spinner struct {
	rows   []*Row
	ticker <-chan time.Time
	stopCh chan struct{}
}

func NewSpinner() *Spinner {
	t := time.NewTicker(200 * time.Millisecond)
	return &Spinner{rows: []*Row{}, ticker: t.C, stopCh: make(chan struct{})}
}

func (s *Spinner) AddRow(startText string) *Row {
	row := &Row{Status: PendingStatus, Message: startText}
	s.rows = append(s.rows, row)
	return row
}

func (s *Spinner) Start() {
	for i := range s.rows {
		fmt.Fprintf(os.Stdout, "%s %s\n", nextTickerChar(), s.rows[i].Message)
	}
	go func() {
		for {
			select {
			case <-s.ticker:
				for i := range s.rows {
					fmt.Fprint(os.Stdout, UpLine)
					fmt.Fprintf(os.Stdout, CursorForward(len(s.rows[i].Message)+padding))
					fmt.Fprintf(os.Stdout, DeleteLine)
				}
				fmt.Fprint(os.Stdout, CursorStartLine)
				tc := nextTickerChar()
				for i := range s.rows {
					fmt.Fprintf(os.Stdout, "%s %s\n", tc, s.rows[i].Message)
				}
			case <-s.stopCh:
				break
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.stopCh <- struct{}{}
}

type Row struct {
	Message string
	Status  Status
}

var tickerChars = []string{"\\", "|", "/", "-"}
var currentTickerChar = 0

func nextTickerChar() string {
	currentTickerChar++
	currentTickerChar %= len(tickerChars)
	return tickerChars[currentTickerChar]
}

func CursorForward(n int) string {
	return fmt.Sprintf("\x1b[%dC", n)
}
