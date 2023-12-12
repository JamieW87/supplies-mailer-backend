package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CreateUserEntryRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Category string `json:"category" binding:"required"`
}

type CreateUserEntryResponse struct {
	UserId string `json:"user-id"`
}

type User struct {
	Id        uuid.UUID `bun:",pk"`
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCategory struct {
	bun.BaseModel `bun:"table:user_categories"`
	UserId        uuid.UUID
	CategoryId    int
}

type Category struct {
	bun.BaseModel `bun:"table:categories"`
	Id            int `bun:",pk"`
	Name          string
	Notes         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
