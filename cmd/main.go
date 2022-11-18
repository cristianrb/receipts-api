package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"receipts-api/internal/handlers"
	"receipts-api/internal/repositories"
	"receipts-api/internal/services"
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
	receiptsRepository := repositories.New(openMysqlConnection())
	receiptsService := services.New(receiptsRepository)
	receiptsHandler := handlers.New(receiptsService)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: receiptsHandler.Routes(),
	}

	println("Started receipts application!")
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func openMysqlConnection() *sql.DB {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
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
