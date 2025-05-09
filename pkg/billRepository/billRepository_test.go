package billRepository

import (
	"biller/mocks"
	"biller/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBill(t *testing.T) {
	dataHandlerMock := new(mocks.BillsJSONStorageHandlerMock)
	repo := NewLocalBillRepository(dataHandlerMock)

	products := []utils.BillProduct{
		{Id: "1", Quantity: 1},
	}
	subtotal := 10.0
	total := 10.0

	newBill := &utils.Bill{
		Products:  products,
		Subtotal:  subtotal,
		Total:     total,
		CreatedAt: "02-01-2006 15:04:05",
		Id:        "1111",
	}

	dataHandlerMock.On("Add", *newBill).Return(newBill, nil).Once()
	result := repo.AddBill(products, subtotal, total, "02-01-2006 15:04:05", "1111")
	assert.Equal(t, newBill, result)

	dataHandlerMock.AssertExpectations(t)
}
