package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
