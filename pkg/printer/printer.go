package printer

import "fmt"

// another printer will be created (for a real printer) and passed to bill

type TerminalPrinter struct{}

func NewTerminalPrinter() *TerminalPrinter {
	return &TerminalPrinter{}
}

func (printer *TerminalPrinter) Print(data string) {
	fmt.Print(data)
}
