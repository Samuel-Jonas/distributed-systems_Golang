package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var errsql error

func OpenConnection(PostgresDriver, ConnectionString, DbName string) (*sql.DB, error) {
	fmt.Printf("Acessing %s... ", DbName)
	db, errsql = sql.Open(PostgresDriver, ConnectionString)

	return db, errsql
}

func CloseConnection(db *sql.DB) {

	errsql = db.Close()
}
