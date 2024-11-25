package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	MemberId int64 `json:"member_id"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type RegisterResponse struct {
	MemberId int64 `json:"member_id"`
}
