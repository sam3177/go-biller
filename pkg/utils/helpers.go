package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

func OpenFileInVsCode(filePath string) {
	cmd := exec.Command("code", filePath) // `code` is the VSCode CLI command
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if error := cmd.Run(); error != nil {
		fmt.Printf("Failed to open file in VSCode: %v", error)
		panic(error)
	}
}

func GetBillsDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to determine caller information")
	}

	// Resolve project root (go up 2 levels from cmd/bill/main.go)
	return filepath.Join(filepath.Dir(file), "../../bills")
}

func CleanBufferBeforeCreatingTheFile(input *bytes.Buffer) *bytes.Buffer {
	var cleaned bytes.Buffer
	for {
		b, err := input.ReadByte()
		if err != nil {
			break // EOF reached
		}
		// Check if the byte is printable or a valid whitespace (newline, tab, etc.)
		if (b >= 32 && b <= 126) || b == '\n' || b == '\t' {
			cleaned.WriteByte(b)
		}
	}

	// Define unwanted sequences using a regex pattern
	// Match sequences like @L, aE!, daE!, E before Total, and dV
	unwantedPattern := regexp.MustCompile(`(@L|aE!|daE!|dV)`)

	// Replace unwanted sequences in the buffer
	content := unwantedPattern.ReplaceAll(cleaned.Bytes(), []byte(""))
	content = regexp.MustCompile(`ETotal`).ReplaceAll(content, []byte("Total"))

	// Return the cleaned content as a bytes.Buffer
	return bytes.NewBuffer(content)
}
