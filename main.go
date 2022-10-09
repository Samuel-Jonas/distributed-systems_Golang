package main

import (
	"database/sql"
	"fmt"

	dbConfig "dataViewer/config"

	Entity "dataViewer/entities"

	database "dataViewer/database"
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
		var Users Entity.User

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

	db, errsql = database.OpenConnection(dbConfig.PostgresDriver, dbConfig.ConnectionString, dbConfig.DbName)

	if errsql != nil {
		panic(errsql.Error())
	} else {
		fmt.Println("Connected!")
	}

	checkErr(errsql)

	defer database.CloseConnection(db)

	sqlSelect()
}
