package user

type PublicUser struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	MCID         string `json:"mcid"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	MCIDVerified bool   `json:"mcidVerified"`
	Admin        bool   `json:"admin"`
}

type UpdateUserRequest struct {
	Value string `json:"value"`
}
