package structs

type Attendance struct {
	Sr             int    `json:"sr" validate:"required"`
	Id             int    `json:"id" validate:"required"`
	Date           string `json:"date" validate:"required"`
	CurrentStatus  string `json:"currentStatus" validate:"required"`
	Intime         string `json:"intime" validate:"required"`
	Outtime        string `json:"outtime" validate:"required"`
	EffectiveHours string `json:"effectiveHours" validate:"required"`
}
