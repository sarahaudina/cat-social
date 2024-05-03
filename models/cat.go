package models

import (
	"time"
)

type Cat struct {
	ID        uint   `json:"id"`
	Name  string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Gender string `json:"gender"`
	UserId uint   `json:"user_id"`
}

type InsertCat struct {
	Name  string `json:"name"`
	Gender string `json:"gender"`
	UserId uint   `json:"user_id"`
}

type InsertCat_ struct {
	Name  string `json:"name"`
	Gender string `json:"gender"`
}