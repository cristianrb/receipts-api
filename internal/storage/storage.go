package storage

import "receipts-api/pkg/types"

type ReceiptStorage interface {
	CreateReceipt(receipt *types.Receipt) (*types.Receipt, error)
	GetReceiptById(receiptId int) (*types.Receipt, error)
	DeleteReceiptById(receiptId int) error
	GetAllReceipts() (types.Receipts, error)
}

type ItemStorage interface {
	CreateItem(item *types.Item) (*types.Item, error)
	GetItems(ids []any) (types.Items, error)
	UpdateItem(item *types.Item) (*types.Item, error)
	DeleteItemById(id int64) error
}
