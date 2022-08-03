package requestTypes

type AdminAuth struct {
	AdminId  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type AddNewDepartment struct {
	Code string `json:"department_code" binding:"required,departmentCodeValidation"`
	Name string `json:"name" binding:"required"`
}

type UpdateDepartment struct {
	Name string `json:"name" binding:"required_without=Code"`
	Code string `json:"department_code" binding:"required_without=Name,omitempty,departmentCodeValidation"`
}

type AddNewStudent struct {
	FirstName    string `json:"first_name" binding:"required,alpha"`
	LastName     string `json:"last_name" binding:"required,alpha"`
	Semester     string `json:"semester" binding:"required,semester"`
	DepartmentId uint   `json:"department_id" binding:"required,numeric"`
	Address      string `json:"address"`
}

type UpdateStudent struct {
	FirstName string `json:"first_name" binding:"omitempty,alpha"`
	LastName  string `json:"last_name" binding:"omitempty,alpha"`
	Address   string `json:"address"`
}

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
