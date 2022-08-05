package responseTypes

type StudentResponse struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	RollNo         string `json:"roll_no"`
	Email          string `json:"email"`
	Semester       string `json:"semester"`
	DepartmentID   uint   `json:"department_id"`
	DepartmentCode string `json:"department_code"`
	DepartmentName string `json:"department_name"`
	Address        string `json:"address,omitempty"`
	ProfileImage   string `json:"profile_image,omitempty"`
}

// type StudentResponse struct {
// 	ID             uint   `gorm:"id" json:"id"`
// 	FirstName      string `gorm:"first_name" json:"first_name"`
// 	LastName       string `gorm:"last_name" json:"last_name"`
// 	RollNo         string `gorm:"roll_no" json:"roll_no"`
// 	Email          string `gorm:"email" json:"email"`
// 	Semester       string `gorm:"semester" json:"semester"`
// 	DepartmentID   uint   `gorm:"department_id" json:"department_id"`
// 	DepartmentCode string `gorm:"department_code" json:"department_code"`
// 	DepartmentName string `gorm:"department_code" json:"department_name"`
// 	Address        string `gorm:"address" json:"address"`
// 	ProfileImage   string `gorm:"profile_image" json:"profile_image"`
// }
