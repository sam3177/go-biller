package fileHandler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

type FileHandler struct {
	FileDir string
}

func NewFileHandler(billsDir string) *FileHandler {
	return &FileHandler{
		FileDir: billsDir,
	}
}

func (handler *FileHandler) Save(data *bytes.Buffer, fileName string) string {

	filePath := handler.FileDir + "/" + fileName

	error := os.WriteFile(filePath, handler.CleanBufferBeforeCreatingTheFile(data).Bytes(), 0644)

	if error != nil {
		fmt.Println("Error", error)
	}

	handler.OpenFile(filePath, "")

	return fileName
}

func (handler *FileHandler) OpenFile(filePath string, command string) {
	if command == "" {
		command = "code" // default command for VSCode (just my choice)
	}
	cmd := exec.Command(command, filePath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if error := cmd.Run(); error != nil {
		fmt.Printf("Failed to open file: %v", error)
		panic(error)
	}
}

func (handler *FileHandler) CleanBufferBeforeCreatingTheFile(input *bytes.Buffer) *bytes.Buffer {
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
	unwantedPattern := regexp.MustCompile(`(@L|aE!|daE!|dV|aE|daE|av0&,)`)

	// Replace unwanted sequences in the buffer
	content := unwantedPattern.ReplaceAll(cleaned.Bytes(), []byte(""))
	content = regexp.MustCompile(`ETotal`).ReplaceAll(content, []byte("Total"))

	// Return the cleaned content as a bytes.Buffer
	return bytes.NewBuffer(content)
}
