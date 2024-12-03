package inputHandler

import (
	"biller/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	value, error := reader.ReadString('\n')

	if error != nil {
		fmt.Println("Error:", error)
	}

	return strings.TrimSpace(value), error
}

func GetValidIntFromInput(reader *bufio.Reader, prompt string, options utils.GetValidNumberFromInputOptions) int {
	value, _ := GetInput(reader, prompt)

	intValue, error := strconv.ParseInt(value, 10, 0)
	if error != nil {
		fmt.Println("Error:", error)
		return GetValidIntFromInput(reader, prompt, options)
	} else if options.ShouldBePositive && intValue <= 0 {
		fmt.Println("Error: You must enter a positive integer number")
		return GetValidIntFromInput(reader, prompt, options)
	}

	return int(intValue)
}

func GetValidFloatFromInput(reader *bufio.Reader, prompt string, options utils.GetValidNumberFromInputOptions) float64 {
	value, _ := GetInput(reader, prompt)

	floatValue, error := strconv.ParseFloat(value, 64)
	if error != nil {
		fmt.Println("Error:", error)
		return GetValidFloatFromInput(reader, prompt, options)
	} else if options.ShouldBePositive && floatValue <= 0 {
		fmt.Println("Error: You must enter a positive decimal number")
		return GetValidFloatFromInput(reader, prompt, options)
	}

	return floatValue
}

func GetTableName() string {
	reader := bufio.NewReader(os.Stdin)
	tableName, error := GetInput(reader, "Please, type the table name: ")

	if error != nil {
		fmt.Println("Error:", error)
		return GetTableName()
	}

	return tableName
}
