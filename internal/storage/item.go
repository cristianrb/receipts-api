package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"receipts-api/internal/logger"
	"receipts-api/pkg/types"
	"strings"
)

var _ ItemStorage = &ItemStorageImpl{}

const (
	queryInsertItem = `INSERT INTO items (product_name) VALUES (?)`
	queryItemById   = `SELECT * from items where id = ?`
	queryUpdateItem = `UPDATE items SET product_name = ? WHERE id = ?`
	queryDeleteItem = `DELETE FROM items WHERE id = ?`
)

type ItemStorageImpl struct {
	Conn *sql.DB
}

// NewItemStorage creates an instance of ItemStorageImpl
func NewItemStorage(conn *sql.DB) *ItemStorageImpl {
	return &ItemStorageImpl{
		Conn: conn,
	}
}

func (s *ItemStorageImpl) CreateItem(item *types.Item) (*types.Item, error) {
	statement, err := s.Conn.Prepare(queryInsertItem)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	insertResult, err := statement.Exec(item.ProductName)
	if err != nil {
		return nil, err
	}

	item.Id, err = insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemStorageImpl) GetItems(ids []any) (types.Items, error) {
	query := `SELECT * from items where id in (?` + strings.Repeat(",?", len(ids)-1) + `)`

	rows, err := s.Conn.Query(query, ids...)
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

func (s *ItemStorageImpl) GetItemById(id int64) (*types.Item, error) {
	row := s.Conn.QueryRow(queryItemById, id)
	item := &types.Item{}
	err := row.Scan(&item.Id, &item.ProductName)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemStorageImpl) UpdateItem(item *types.Item) (*types.Item, error) {
	_, err := s.GetItemById(item.Id)
	if err != nil {
		logger.Error(fmt.Sprintf("Retrieving item with id: %d to update but does not exist", item.Id), err)
		return nil, err
	}

	res, err := s.Conn.Exec(queryUpdateItem, item.ProductName, item.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("Item with id: %d was not updated", item.Id))
	}

	return item, nil
}

func (s *ItemStorageImpl) DeleteItemById(id int64) error {
	res, err := s.Conn.Exec(queryDeleteItem, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("Item with id: %d was not deleted", id))
	}

	return nil
}
