package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	var err error

	DB, err := sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	err = createTables(DB)
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")

	return DB
}

func createTables(db *sql.DB) error {

	createUsersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        )
    `
	_, err := db.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createConferencesTable := `
        CREATE TABLE IF NOT EXISTS conferences (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
            FOREIGN KEY(user_id) REFERENCES users(id)
        )
    `

	_, err = db.Exec(createConferencesTable)
	if err != nil {
		panic("Could not create conferences table.")
	}

	return err
}
