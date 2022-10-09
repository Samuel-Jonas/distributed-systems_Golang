package dbconfig

import "fmt"

const (
	User     = "postgres"
	Host     = "localhost"
	Port     = "5432"
	Password = "LOSfro@#33"
	Name     = "sd_servicesystem"
)

var ConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	Host,
	Port,
	User,
	Name,
	Password,
)
