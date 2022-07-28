package responses

type AdminAuthSuccess struct {
	Base
	Data TokenResponse `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Base
}

type AuthError struct {
	Base
}
