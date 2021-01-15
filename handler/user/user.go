package user

type GetUserInfoRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type GetUserInfoResponse struct {
	Uid int64 `json:"uid"`
}

