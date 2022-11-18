package repositories

import (
	"database/sql"
	"receipts-api/pkg/models"
	"time"
)

var _ ReceiptsRepository = &ReceiptsRepositoryImpl{}

type ReceiptsRepository interface {
	CreateReceipt(receipt *models.Receipt) (*models.Receipt, error)
}

type ReceiptsRepositoryImpl struct {
	Conn *sql.DB
}

// New creates an instance of ReceiptsRepositoryImpl
func New(conn *sql.DB) *ReceiptsRepositoryImpl {
	return &ReceiptsRepositoryImpl{
		Conn: conn,
	}
}

// CreateReceipt creates a receipt in the database
func (r *ReceiptsRepositoryImpl) CreateReceipt(receipt *models.Receipt) (*models.Receipt, error) {
	receipt.CreatedOn = time.Now()

	var itemsInsertErr error
	receipt.Items, itemsInsertErr = r.insertItems(receipt.Items)
	if itemsInsertErr != nil {
		return nil, itemsInsertErr
	}

	receipt, receiptInsertErr := r.insertReceipt(receipt)
	if receiptInsertErr != nil {
		return nil, receiptInsertErr
	}

	receiptProductsErr := r.insertReceiptProducts(receipt)
	if receiptProductsErr != nil {
		return nil, receiptProductsErr
	}

	return receipt, nil
}

// insertItems inserts a list of items in the database
func (r *ReceiptsRepositoryImpl) insertItems(items models.Items) (models.Items, error) {
	query := `INSERT INTO items (product_name) VALUES (?);`
	statement, err := r.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	for _, item := range items {
		insertResult, err := statement.Exec(item.ProductName)
		if err != nil {
			return nil, err
		}

		item.Id, err = insertResult.LastInsertId()
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

// insertReceipt inserts a receipt in the database
func (r *ReceiptsRepositoryImpl) insertReceipt(receipt *models.Receipt) (*models.Receipt, error) {
	query := `INSERT INTO receipts (created_on) VALUES (?);`
	statement, err := r.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	insertResult, err := statement.Exec(receipt.CreatedOn)
	if err != nil {
		return nil, err
	}

	receipt.Id, err = insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

// insertReceiptProducts insert the relationship between receipt and items in the database
func (r *ReceiptsRepositoryImpl) insertReceiptProducts(receipt *models.Receipt) error {
	query := `INSERT INTO receipt_product (receipt_id, product_id) VALUES (?, ?);`
	statement, err := r.Conn.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	for _, item := range receipt.Items {
		_, err := statement.Exec(receipt.Id, item.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
