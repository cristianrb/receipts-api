package storage

import "receipts-api/pkg/types"

type Storage interface {
	CreateReceipt(receipt *types.Receipt) (*types.Receipt, error)
	GetReceiptById(receiptId int) (*types.Receipt, error)
	GetAllReceipts() (types.Receipts, error)
}
