package responseTypes

type Base struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Meta struct {
	Count int `json:"count"`
	Base
}

type SuccessResponseData struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
