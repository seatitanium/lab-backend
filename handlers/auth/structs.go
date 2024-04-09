package auth

// 前端发来的用于注册的内容
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	MCID     string `json:"mcid"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
