package types

type RefreshRequestObject struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginRequestObject struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
