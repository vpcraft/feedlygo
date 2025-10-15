package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/vpcraft/feedlygo/internal/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fullname  string    `json:"fullname"`
}

func serializerUser(user db.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Fullname:  user.Fullname,
	}
}
