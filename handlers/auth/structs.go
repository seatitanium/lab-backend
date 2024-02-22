package auth

import "time"

type User struct {
	Id        int64     `json:"id" db:"id"`
	Nickname  string    `json:"nickname" db:"nickname"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	MCID      string    `json:"mcid" db:"mcid"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Hash      string    `db:"hash"`
}

// 前端发来的用于注册的内容
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	MCID     string `json:"mcid"`
}
