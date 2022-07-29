package requestTypes

type AdminAuth struct {
	AdminId  string `json:"adminId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminChangePassword struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type AddNewDepartment struct {
	Code string `json:"department_code" binding:"required,departmentCodeValidation"`
	Name string `json:"" binding:"required"`
}
