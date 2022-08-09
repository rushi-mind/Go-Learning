package requestTypes

type AddAssignment struct {
	Name         string `json:"name" binding:"required"`
	Semester     string `json:"semester" binding:"required,semesterValidation"`
	DepartmentID uint   `json:"department_id" binding:"required,numeric"`
	Deadline     string `json:"deadline" binding:"required"`
}
