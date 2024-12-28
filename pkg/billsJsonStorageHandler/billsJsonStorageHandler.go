package billsJsonStorageHandler

import (
	"biller/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
)

type BillsJSONStorageHandler struct {
	FilePath string
}

func NewBillsJSONStorageHandler(filePath string) *BillsJSONStorageHandler {
	return &BillsJSONStorageHandler{
		FilePath: filePath,
	}
}

func (handler *BillsJSONStorageHandler) GetAll() ([]utils.Bill, error) {
	file, err := os.Open(handler.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var bills []utils.Bill
	if err := json.Unmarshal(bytes, &bills); err != nil {
		log.Fatal("failed to parse JSON: %w", err)
	}

	return bills, nil
}

func (handler *BillsJSONStorageHandler) Get(id string) (*utils.Bill, error) {
	bills, err := handler.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read bills: %w", err)
	}

	for _, bill := range bills {
		if bill.Id == id {
			return &bill, nil
		}
	}

	return nil, fmt.Errorf("bill with ID %v not found", id)
}

func (handler *BillsJSONStorageHandler) Add(newBill utils.Bill) (*utils.Bill, error) {
	bills, err := handler.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read bills: %w", err)
	}

	newBill.Id = uuid.NewString()

	bills = append(bills, newBill)

	file, err := os.OpenFile(handler.FilePath, os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := json.Marshal(bills)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bills: %w", err)
	}

	_, err = file.Write(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to write to file: %w", err)
	}

	return &newBill, nil
}
