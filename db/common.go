package db

import "database/sql"

func OpenDB() *sql.DB {
	db, _ := sql.Open("sqlite3", "./bdd.db")

	return db
}