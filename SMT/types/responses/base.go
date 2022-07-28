package responses

type Base struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
