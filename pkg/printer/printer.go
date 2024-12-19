package printer

import (
	"bytes"
	"fmt"
)

// another printer will be created (for a real printer) and passed to bill

type TerminalPrinter struct {
	rowLength int
}

func NewTerminalPrinter(rowLength int) *TerminalPrinter {
	return &TerminalPrinter{
		rowLength: rowLength,
	}
}

func (printer *TerminalPrinter) Print(data bytes.Buffer) {
	fmt.Print(data.String())
}

func (printer *TerminalPrinter) GetRowLength() int {
	return printer.rowLength
}
