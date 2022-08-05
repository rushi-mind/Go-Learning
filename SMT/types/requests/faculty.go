package requestTypes

type AddNewTeacher struct {
	FirstName    string `json:"first_name" binding:"required,alpha"`
	LastName     string `json:"last_name" binding:"required,alpha"`
	DepartmentId uint   `json:"department_id" binding:"required,numeric"`
}

type UpdateTeacher struct {
	FirstName    string `json:"first_name" binding:"alpha"`
	LastName     string `json:"last_name" binding:"alpha"`
	DepartmentId uint   `json:"department_id" binding:"numeric"`
}
