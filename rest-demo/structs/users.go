package structs

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=12"`
	Pin  string `json:"pin" validate:"required,min=4,max=4"`
}
