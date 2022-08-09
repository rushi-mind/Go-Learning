package models

type Assignment struct {
	ID           uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string     `gorm:"column:name;type:varchar(50)" json:"name"`
	Semester     string     `gorm:"column:semester;type:enum('1','2','3','4','5','6','7','8')" json:"semester"`
	DepartmentID uint       `gorm:"column:department_id;type:uint" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"-"`
	Deadline     string     `gorm:"column:deadline;type:date" json:"deadline"`
	FilePath     string     `gorm:"column:file_path;type:varchar(200)" json:"file_path"`
}

func (Assignment) TableName() string {
	return "assignments"
}
