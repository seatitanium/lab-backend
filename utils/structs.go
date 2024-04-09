package utils

import (
	"database/sql"
	"time"
)

type JWTPayload struct {
	Username  string `json:"username"`
	UpdatedAt int64  `json:"updated_at"`
}

type DbInstance struct {
	Id           uint      `json:"id" db:"id"`
	InstanceId   string    `json:"instance_id" db:"instance_id"`
	TradePrice   float32   `json:"trade_price" db:"trade_price"`
	RegionId     string    `json:"region_id" db:"region_id"`
	InstanceType string    `json:"instance_type" db:"instance_type"`
	Active       bool      `json:"active" db:"active"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type DbUser struct {
	Id        uint           `json:"id" db:"id"`
	Nickname  sql.NullString `json:"nickname" db:"nickname"`
	Username  string         `json:"username" db:"username"`
	Email     string         `json:"email" db:"email"`
	MCID      string         `json:"mcid" db:"mcid"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	Hash      string         `db:"hash"`
}
