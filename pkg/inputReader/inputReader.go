package inputReader

import (
	"bufio"
	"fmt"
	"strings"
)

type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader(reader *bufio.Reader) *InputReader {
	return &InputReader{reader: reader}
}

func (ir *InputReader) GetInput(prompt string) (string, error) {
	fmt.Print(prompt)
	value, error := ir.reader.ReadString('\n')

	if error != nil {
		fmt.Println("Error:", error)
	}

	return strings.TrimSpace(value), error
}
