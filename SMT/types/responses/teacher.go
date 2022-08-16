package responseTypes

type TeacherResponse struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	EmployeeID     string `json:"employee_id"`
	Email          string `json:"email"`
	DepartmentID   uint   `json:"department_id"`
	DepartmentCode string `json:"department_code"`
	DepartmentName string `json:"department_name"`
}
