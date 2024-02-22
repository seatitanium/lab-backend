package utils

type JWTPayload struct {
	Username  string `json:"username"`
	UpdatedAt int64  `json:"updated_at"`
}
