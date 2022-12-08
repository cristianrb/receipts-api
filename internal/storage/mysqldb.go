package storage

import (
	"database/sql"
	"receipts-api/pkg/types"
	"time"
)

var _ Storage = &MysqlStorage{}

type MysqlStorage struct {
	Conn *sql.DB
}

// New creates an instance of MysqlStorage
func New(conn *sql.DB) *MysqlStorage {
	return &MysqlStorage{
		Conn: conn,
	}
}

// CreateReceipt creates a receipt in the database
func (s *MysqlStorage) CreateReceipt(receipt *types.Receipt) (*types.Receipt, error) {
	receipt.CreatedOn = time.Now()

	var itemsInsertErr error
	receipt.Items, itemsInsertErr = s.insertItems(receipt.Items)
	if itemsInsertErr != nil {
		return nil, itemsInsertErr
	}

	receipt, receiptInsertErr := s.insertReceipt(receipt)
	if receiptInsertErr != nil {
		return nil, receiptInsertErr
	}

	receiptProductsErr := s.insertReceiptProducts(receipt)
	if receiptProductsErr != nil {
		return nil, receiptProductsErr
	}

	return receipt, nil
}

func (s *MysqlStorage) GetReceiptById(receiptId int) (*types.Receipt, error) {
	receipt, err := s.findReceiptById(receiptId)
	if err != nil {
		return nil, err
	}
	receiptItems, err := s.findReceiptItemsByReceiptId(receiptId)
	if err != nil {
		return nil, err
	}

	receipt.Items = receiptItems
	return receipt, nil
}

func (s *MysqlStorage) GetAllReceipts() (types.Receipts, error) {
	query := `SELECT r.id AS receipt_id, r.created_on, i.id AS product_id, i.product_name
			  FROM receipts r
			  INNER JOIN receipt_product rp
			  ON rp.receipt_id = r.id
			  INNER JOIN items i
			  ON i.id = rp.product_id`

	statement, err := s.Conn.Prepare(query)
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

// insertItems inserts a list of items in the database
func (s *MysqlStorage) insertItems(items types.Items) (types.Items, error) {
	query := `INSERT INTO items (product_name) VALUES (?);`
	statement, err := s.Conn.Prepare(query)
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
func (s *MysqlStorage) insertReceipt(receipt *types.Receipt) (*types.Receipt, error) {
	query := `INSERT INTO receipts (created_on) VALUES (?);`
	statement, err := s.Conn.Prepare(query)
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
func (s *MysqlStorage) insertReceiptProducts(receipt *types.Receipt) error {
	query := `INSERT INTO receipt_product (receipt_id, product_id) VALUES (?, ?);`
	statement, err := s.Conn.Prepare(query)
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

func (s *MysqlStorage) findReceiptById(id int) (*types.Receipt, error) {
	query := `SELECT id, created_on FROM db.receipts WHERE id = ?`
	statement, err := s.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := s.Conn.QueryRow(query, id)
	receipt := new(types.Receipt)
	err = row.Scan(&receipt.Id, &receipt.CreatedOn)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (s *MysqlStorage) findReceiptItemsByReceiptId(id int) (types.Items, error) {
	query := `SELECT i.id, i.product_name FROM db.items i, db.receipt_product r WHERE receipt_id = ? AND i.id = r.product_id;`
	statement, err := s.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(id)
	if err != nil {
		return nil, err
	}

	items := types.Items{}
	for rows.Next() {
		item := types.Item{}
		err := rows.Scan(&item.Id, &item.ProductName)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
