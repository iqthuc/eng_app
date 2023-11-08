package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConfigDatabase() (*sql.DB, error) {
	username := "root"
	password := "root"
	host := "localhost:3306"
	database := "english_app_db"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
