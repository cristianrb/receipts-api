package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"receipts-api/internal/api"
	"receipts-api/internal/storage"
)

const PORT = "8080"

const (
	mysqlUsername = "MYSQL_USER"
	mysqlPassword = "MYSQL_PASSWORD"
	mysqlHost     = "MYSQL_HOST"
	mysqlSchema   = "MYSQL_DATABASE"
)

var (
	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
	host     = os.Getenv(mysqlHost)
	schema   = os.Getenv(mysqlSchema)
)

func main() {
	mysqlDB := storage.New(openMysqlConnection())
	server := api.New(fmt.Sprintf(":%s", PORT), mysqlDB)
	server.Run()
}

func openMysqlConnection() *sql.DB {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		username, password, host, schema,
	)
	var err error
	client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}

	return client
}
