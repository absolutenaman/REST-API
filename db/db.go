package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "api.DB")
	if err != nil {
		panic("DB connection successful")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()

}
func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS events  (
     id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL ,
    description TEXT NOT NULL ,
    location    TEXT NOT NULL ,
    dateTime    DATETIME NOT NULL ,
    userId      INTEGER 
)`
	_, err := DB.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

}
