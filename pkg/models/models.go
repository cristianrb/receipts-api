package models

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
