package requestTypes

type AddNewStudent struct {
	FirstName    string `json:"first_name" binding:"required,alpha"`
	LastName     string `json:"last_name" binding:"required,alpha"`
	Semester     string `json:"semester" binding:"required,semesterValidation"`
	DepartmentId uint   `json:"department_id" binding:"required,numeric"`
	Address      string `json:"address"`
}

type UpdateStudent struct {
	FirstName string `json:"first_name" binding:"omitempty,alpha"`
	LastName  string `json:"last_name" binding:"omitempty,alpha"`
	Address   string `json:"address"`
}
