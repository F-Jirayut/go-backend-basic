package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver   = "mysql"
	dbUsername = "root"
	dbPassword = ""
	dbName     = "go_basic"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
