package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func initializeDatabase() {
	username := "root"
	password := "root"
	host := "localhost:3306"
	database := "english_app_db"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, database)
	newDb, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = newDb.PingContext(ctx)
	if err != nil {
		log.Println(err)

	}
	log.Println("Connected to MySQL database.")
	db = newDb
}

func GetDB() *sql.DB {
	dbOnce.Do(initializeDatabase)
	return db
}
func ReconnectMySql() {
	if db == nil {
		log.Println("Database connection is not initialized.")
		return
	}

	err := db.Ping()
	if err != nil {
		log.Println("Lost connection to database. Reconnecting...")
		initializeDatabase()
	}
}
func CloseMysqlConnection() {
	if db == nil {
		return
	}
	err := db.Close()

	if err != nil {
		log.Println(err)
	}
	log.Println("Connection to MySql closed.")
}
