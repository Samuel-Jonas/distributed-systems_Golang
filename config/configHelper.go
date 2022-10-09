package dbconfig

import "fmt"

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "LOSfro@#33"
const DbName = "sd_servicesystem"
const TableName = "users"

var ConnectionString = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
