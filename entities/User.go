package entities

type User struct {
	UserId       int    `json:"userId"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	CreationDate string `json:"creationdate"`
}
