package models

type Faculty struct {
	Id           uint       `gorm:"primaryKey;autoIncrement;column:id;type:uint" json:"id"`
	FirstName    string     `gorm:"column:first_name;type:varchar(50)" json:"first_name"`
	LastName     string     `gorm:"column:last_name;type:varchar(50)" json:"last_name"`
	EmployeeId   string     `gorm:"column:employee_id;unique;type:varchar(9)" json:"employee_id"`
	Email        string     `gorm:"column:email;type:varchar(50);unique" json:"email"`
	Password     string     `gorm:"column:password;type:varchar(100);" json:"password"`
	DepartmentId uint       `gorm:"column:department_id;type:uint" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentId" json:"-"`
}

func (Faculty) TableName() string {
	return "faculties"
}
