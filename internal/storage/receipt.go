package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"receipts-api/pkg/types"
	"time"
)

var _ ReceiptStorage = &ReceiptStorageImpl{}

const (
	queryInsertReceipt        = `INSERT INTO receipts (created_on) VALUES (?);`
	queryInsertReceiptProduct = `INSERT INTO receipt_product (receipt_id, product_id) VALUES (?, ?);`
	queryFindAllReceipts      = `SELECT r.id AS receipt_id, r.created_on, i.id AS product_id, i.product_name
			  FROM receipts r
			  INNER JOIN receipt_product rp
			  ON rp.receipt_id = r.id
			  INNER JOIN items i
			  ON i.id = rp.product_id`
	queryFindReceiptById = `SELECT r.id AS receipt_id, r.created_on, i.id AS product_id, i.product_name
			  FROM receipts r
			  INNER JOIN receipt_product rp
			  ON rp.receipt_id = r.id
			  INNER JOIN items i
			  ON i.id = rp.product_id
			  WHERE r.id = ?`
	queryDeleteReceiptById     = `DELETE FROM receipts WHERE id = ?`
	queryDeleteReceiptProducts = `DELETE FROM receipt_product WHERE receipt_id = ?`
)

type ReceiptStorageImpl struct {
	Conn         *sql.DB
	ItemsStorage ItemStorage
}

// NewReceiptStorage creates an instance of ReceiptStorageImpl
func NewReceiptStorage(conn *sql.DB, itemsStorage ItemStorage) *ReceiptStorageImpl {
	return &ReceiptStorageImpl{
		Conn:         conn,
		ItemsStorage: itemsStorage,
	}
}

// CreateReceipt creates a receipt in the database
func (s *ReceiptStorageImpl) CreateReceipt(receiptReq *types.ReceiptRequest) (*types.Receipt, error) {
	receipt := &types.Receipt{
		CreatedOn: time.Now(),
	}
	itemsAsAny := []any{}
	for _, item := range receiptReq.Items {
		itemsAsAny = append(itemsAsAny, item)
	}
	items, err := s.ItemsStorage.GetItems(itemsAsAny)
	if err != nil {
		return nil, err
	}
	if len(items) != len(receiptReq.Items) {
		return nil, errors.New("some of the items does not exist")
	}

	receipt.Items = items
	receipt, receiptInsertErr := s.insertReceipt(receipt)
	if receiptInsertErr != nil {
		return nil, receiptInsertErr
	}

	receiptProductsErr := s.insertReceiptProducts(receipt, nil)
	if receiptProductsErr != nil {
		return nil, receiptProductsErr
	}

	return receipt, nil
}

func (s *ReceiptStorageImpl) GetReceiptById(receiptId int) (*types.Receipt, error) {
	statement, err := s.Conn.Prepare(queryFindReceiptById)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(receiptId)
	if err != nil {
		return nil, err
	}

	rwps := types.ReceiptsWithItemsFromDB{}
	for rows.Next() {
		rwp := types.ReceiptWithItemsFromDB{}
		err := rows.Scan(&rwp.ReceiptId, &rwp.CreatedOn, &rwp.ProductId, &rwp.ProductName)
		if err != nil {
			return nil, err
		}
		rwps = append(rwps, rwp)
	}
	if err != nil {
		return nil, err
	}

	receiptsList := rwps.ToReceipts()
	if len(receiptsList) == 0 {
		return nil, errors.New(fmt.Sprintf("Receipt with id: %d not found", receiptId))
	}
	return &receiptsList[0], nil
}

func (s *ReceiptStorageImpl) GetAllReceipts() (types.Receipts, error) {
	statement, err := s.Conn.Prepare(queryFindAllReceipts)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	rwps := types.ReceiptsWithItemsFromDB{}
	for rows.Next() {
		rwp := types.ReceiptWithItemsFromDB{}
		err := rows.Scan(&rwp.ReceiptId, &rwp.CreatedOn, &rwp.ProductId, &rwp.ProductName)
		if err != nil {
			return nil, err
		}
		rwps = append(rwps, rwp)
	}

	return rwps.ToReceipts(), nil
}

func (s *ReceiptStorageImpl) DeleteReceiptById(receiptId int) error {
	statement, err := s.Conn.Prepare(queryDeleteReceiptById)
	if err != nil {
		return err
	}
	defer statement.Close()

	res, err := statement.Exec(receiptId)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("No receipt to delete with id: %d", receiptId))
	}

	return nil
}

func (s *ReceiptStorageImpl) UpdateReceipt(receiptReq *types.ReceiptRequest) (*types.Receipt, error) {
	tx, err := s.Conn.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	receiptDB, err := s.GetReceiptById(int(receiptReq.Id))
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(queryDeleteReceiptProducts, receiptReq.Id)
	if err != nil {
		return nil, err
	}

	itemsAsAny := []any{}
	for _, item := range receiptReq.Items {
		itemsAsAny = append(itemsAsAny, item)
	}
	items, err := s.ItemsStorage.GetItems(itemsAsAny)
	if err != nil {
		return nil, err
	}
	if len(items) != len(receiptReq.Items) {
		return nil, errors.New("some of the items do not exist")
	}

	receipt := &types.Receipt{
		Id:        receiptReq.Id,
		Items:     items,
		CreatedOn: receiptDB.CreatedOn,
	}
	err = s.insertReceiptProducts(receipt, tx)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return receipt, nil
}

// insertReceipt inserts a receipt in the database
func (s *ReceiptStorageImpl) insertReceipt(receipt *types.Receipt) (*types.Receipt, error) {
	statement, err := s.Conn.Prepare(queryInsertReceipt)
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
func (s *ReceiptStorageImpl) insertReceiptProducts(receipt *types.Receipt, tx *sql.Tx) error {
	if tx == nil {
		tx, err := s.Conn.BeginTx(context.Background(), nil)
		if err != nil {
			return err
		}
		// Defer a rollback in case anything fails.
		defer tx.Rollback()
	}
	statement, err := tx.Prepare(queryInsertReceiptProduct)
	if err != nil {
		return err
	}
	defer statement.Close()

	for _, item := range receipt.Items {
		_, err := statement.Query(receipt.Id, item.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
