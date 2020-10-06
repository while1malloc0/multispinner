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

func main() {
	fmt.Fprint(os.Stdout, "Hello there\n")
	fmt.Fprint(os.Stdout, "An again\n")
	fmt.Fprintf(os.Stdout, "One more")
	time.Sleep(3 * time.Second)
	fmt.Fprintf(os.Stdout, UpLine)
	fmt.Fprintf(os.Stdout, UpLine)
	fmt.Fprintf(os.Stdout, CursorStartLine)
	fmt.Fprint(os.Stdout, "Multi clear\n")
	fmt.Fprint(os.Stdout, "Multi clear\n")
	fmt.Fprint(os.Stdout, "Multi clear\n")
}
