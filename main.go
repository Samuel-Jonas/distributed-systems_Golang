package main

import (
	"database/sql"
	"fmt"

	dbConfig "databasePostgres/dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var errsql error

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func sqlSelect() {
	sqlStatement, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for sqlStatement.Next() {
		var Users dbConfig.Userr

		err := sqlStatement.Scan(
			&Users.UserId,
			&Users.Name,
			&Users.Password,
			&Users.Email,
			&Users.Address,
			&Users.CreationDate,
		)
		checkErr(err)

		fmt.Printf("UserId: %d, Name: %s, Password: %s, Email: %s, Address: %s, CreationDate: %s", Users.UserId, Users.Name, Users.Password, Users.Email, Users.Address, Users.CreationDate)
	}
}

func main() {
	fmt.Printf("Acessing %s... ", dbConfig.DbName)
	db, errsql = sql.Open(dbConfig.PostgresDriver, dbConfig.ConnectionString)

	if errsql != nil {
		panic(errsql.Error())
	} else {
		fmt.Println("Connected!")
	}

	checkErr(errsql)

	defer db.Close()

	sqlSelect()
}
