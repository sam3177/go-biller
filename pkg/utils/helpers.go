package utils

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

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

func RoundToGivenDecimals(value float64, decimals int) float64 {
	precision := math.Pow(10, float64(decimals))

	return math.Round(value*precision) / precision
}
