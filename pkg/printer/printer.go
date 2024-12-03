package printer

import "fmt"

type TerminalPrinter struct{}

func NewTerminalPrinter() *TerminalPrinter {
	return &TerminalPrinter{}
}

func (printer *TerminalPrinter) Print(data string) {
	fmt.Print(data)
}
