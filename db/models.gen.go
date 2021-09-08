// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Item struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Unit      string    `json:"unit"`
	Cost      string    `json:"cost"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
