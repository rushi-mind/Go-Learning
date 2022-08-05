package requestTypes

type AddNewDepartment struct {
	Code string `json:"department_code" binding:"required,departmentCodeValidation"`
	Name string `json:"name" binding:"required"`
}

type UpdateDepartment struct {
	Name string `json:"name" binding:"required_without=Code"`
	Code string `json:"department_code" binding:"required_without=Name,omitempty,departmentCodeValidation"`
}
