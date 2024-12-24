package printer

import (
	"bytes"
	"log"
	"os/exec"
)

type EpsonPrinter struct {
	name      string
	rowLength int
}

func NewEpsonPrinter(printerName string) *EpsonPrinter {
	return &EpsonPrinter{
		name:      printerName,
		rowLength: 46, // hardcoded on purpose, it will not change for this printer
	}
}

func (printer *EpsonPrinter) Print(data bytes.Buffer) {
	cmd := exec.Command("lp", "-d", printer.name, "-o", "raw")

	cmd.Stdin = &data
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error printing: %v", err)
	}
}

func (printer *EpsonPrinter) GetRowLength() int {
	return printer.rowLength
}
