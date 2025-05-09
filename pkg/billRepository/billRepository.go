package billRepository

import (
	"biller/pkg/utils"
	"fmt"
)

type LocalBillRepository struct {
	dataHandler utils.BillStorageHandlerInterface
}

func NewLocalBillRepository(
	dataHandler utils.BillStorageHandlerInterface,
) *LocalBillRepository {
	return &LocalBillRepository{
		dataHandler: dataHandler,
	}
}

func (repo *LocalBillRepository) GetBills() []utils.Bill {
	bills, error := repo.dataHandler.GetAll()

	if error != nil {
		return nil
	}

	return bills
}

func (repo *LocalBillRepository) GetBillById(id string) (*utils.Bill, error) {
	bill, error := repo.dataHandler.Get(id)

	if error != nil {
		return nil, error
	}

	return bill, nil
}

func (repo *LocalBillRepository) AddBill(
	products []utils.BillProduct,
	subtotal float64,
	total float64,
	createdAt string,
	id string,
) *utils.Bill {
	newBill, error := repo.dataHandler.Add(utils.Bill{
		Products:  products,
		Subtotal:  subtotal,
		Total:     total,
		CreatedAt: createdAt,
		Id:        id,
	})

	if error != nil {
		fmt.Println(error)
		return nil
	}

	return newBill
}
