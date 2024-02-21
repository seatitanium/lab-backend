package auth

import "time"

type User struct {
	Id        int64     `json:"id"`
	Nickname  string    `json:"nickname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	MCID      string    `json:"mcid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Hash      string
}

// 前端发来的用于注册的内容
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	MCID     string `json:"mcid"`
}
