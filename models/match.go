package models

import (
	"time"
)

type Match struct {
	ID        uint   `json:"id"`
	Status  string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Issuer uint   `json:"issuer"`
	Receiver uint   `json:"receiver"`
}

type InsertMatch struct {
	Status  string `json:"status"`
	Issuer uint   `json:"issuer"`
	Receiver uint   `json:"receiver"`
}

type InsertMatch_ struct {
	Issuer uint   `json:"issuer"`
	Receiver uint   `json:"receiver"`
}

type UpdateMatch struct {
	ID        uint   `json:"id"`
	Status  string `json:"status"`
}

type DeleteMatch struct {
	ID        uint   `json:"id"`
}

type ResponseToMatchRequest struct {
	ID        uint   `json:"id"`
}