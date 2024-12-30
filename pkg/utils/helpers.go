package utils

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
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

func GetAbsolutePath(relativePath string) string {
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Check if running from a temp directory
	if strings.Contains(executablePath, "go-build") {
		// Assume development mode; use the working directory
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get working directory: %v", err)
		}
		return filepath.Join(wd, relativePath)
	}

	return filepath.Join(filepath.Dir(executablePath), "..", relativePath)
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
	// Match sequences like @L, aE!, daE! and dV
	unwantedPattern := regexp.MustCompile(`(@L|aE!|daE!|dV)`)

	// Replace unwanted sequences in the buffer
	content := unwantedPattern.ReplaceAll(cleaned.Bytes(), []byte(""))
	content = regexp.MustCompile(`ETotal`).ReplaceAll(content, []byte("Total"))

	// Return the cleaned content as a bytes.Buffer
	return bytes.NewBuffer(content)
}

func RoundToGivenDecimals(value float64, decimals int) float64 {
	precision := math.Pow(10, float64(decimals))

	return math.Round(value*precision) / precision
}
