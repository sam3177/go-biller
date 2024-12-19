package printer

import (
	"bytes"
	"log"
	"os/exec"
)

type EpsonPrinter struct {
	name string
}

func NewEpsonPrinter(printerName string) *EpsonPrinter {
	return &EpsonPrinter{
		name: printerName,
	}
}

func (printer *EpsonPrinter) Print(data string) {
	cmd := exec.Command("lp", "-d", printer.name, "-o", "raw")

	cmd.Stdin = printer.getContentBuffer(data)
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error printing: %v", err)
	}
}

func (printer *EpsonPrinter) getContentBuffer(data string) *bytes.Buffer {
	var buffer bytes.Buffer

	buffer.Write([]byte{0x1B, 0x40}) // ESC @ (initialize)

	buffer.Write([]byte{0x1D, 0x4C, 0x0A, 0x00}) // 10 dots left margin (1 character width)

	buffer.Write([]byte{0x1B, 0x61, 0x00}) // Left align

	buffer.WriteString(data)

	buffer.Write([]byte{0x1B, 0x64, 0x03}) // Feed 3 lines
	buffer.Write([]byte{0x1D, 0x56, 0x00}) // Cut the paper

	return &buffer
}
