package ttomato

import (
	"fmt"
	"log"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

var (
	printWidth uint
)

func init() {
	terminalWidth, err := terminal.Width()
	if err != nil {
		log.Fatal(err)
	}
	printWidth = terminalWidth / 2
}

func setBoundary() error {
	fmt.Print("|")
	for i := 0; i < int(printWidth)+1; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|")
	for i := 0; i < int(printWidth)+2; i++ {
		fmt.Printf("\b")
	}
	return nil
}

func printOneTomato(lastone bool) {
	fmt.Print("🍅")
	if !lastone {
		fmt.Print(" ")
	}
}
