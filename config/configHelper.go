package dbconfig

import "fmt"

const (
	User     = "postgres"
	Host     = "database"
	Password = "postgres"
	Name     = "postgres"
)

var ConnectionString = fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable",
	User,
	Password,
	Host,
	Name,
)
