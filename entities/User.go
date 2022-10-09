package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey; unique;"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	CreationDate time.Time `gorm:"autoCreateTime"`
	UpdateDate   time.Time `gorm:"autoCreateTime"`
}
