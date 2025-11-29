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
	createUsersTable := `CREATE TABLE IF NOT EXISTS users  (
     id          INTEGER PRIMARY KEY AUTOINCREMENT,
    email        TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL 
)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS events  (
     id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL ,
    description TEXT NOT NULL ,
    location    TEXT NOT NULL ,
    dateTime    DATETIME NOT NULL ,
    userId INTEGER NOT NULL ,
    FOREIGN KEY(userId) REFERENCES users(id)
)`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	createRegistrationsQuery := `CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT ,
    userId INTEGER NOT NULL ,
    eventId INTEGER NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id) ,
    FOREIGN KEY (eventId) REFERENCES events(id)
    );`
	_, err = DB.Exec(createRegistrationsQuery)
	if err != nil {
		panic(err)
	}
}
