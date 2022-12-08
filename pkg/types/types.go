package types

import "time"

type Item struct {
	Id          int64  `json:"id"`
	ProductName string `json:"product_name"`
}

type Items []*Item

type Receipt struct {
	Id        int64     `json:"id"`
	Items     Items     `json:"items"`
	CreatedOn time.Time `json:"created_on"`
}

type Receipts []Receipt

type ReceiptWithItemsFromDB struct {
	ReceiptId   int64     `json:"receipt_id"`
	CreatedOn   time.Time `json:"created_on"`
	ProductId   int64     `json:"product_id"`
	ProductName string    `json:"product_name"`
}

type ReceiptsWithItemsFromDB []ReceiptWithItemsFromDB

func (rwis ReceiptsWithItemsFromDB) ToReceipts() Receipts {
	receipts := Receipts{}

	if len(rwis) == 0 {
		return receipts
	}

	receipt := Receipt{
		Id:        rwis[0].ReceiptId,
		CreatedOn: rwis[0].CreatedOn,
		Items:     Items{},
	}
	for _, rwp := range rwis {
		if receipt.Id != rwp.ReceiptId {
			receipts = append(receipts, receipt)
			receipt = Receipt{
				Id:        rwp.ReceiptId,
				CreatedOn: rwp.CreatedOn,
				Items:     Items{},
			}
		} else {
			receipt.Items = append(receipt.Items, &Item{ProductName: rwp.ProductName})
		}
	}

	return receipts
}
