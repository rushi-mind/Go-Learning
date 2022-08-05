package responseTypes

type Base struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseData struct {
	Base
	Data any `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
