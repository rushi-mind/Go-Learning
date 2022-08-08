package models

type Student struct {
	Id           uint       `gorm:"primaryKey;autoIncrement;column:id;type:uint" json:"id"`
	RollNo       string     `gorm:"column:roll_no;unique;type:varchar(13)" json:"roll_no"`
	FirstName    string     `gorm:"column:first_name;type:varchar(50)" json:"first_name"`
	LastName     string     `gorm:"column:last_name;type:varchar(50)" json:"last_name"`
	Email        string     `gorm:"unique;column:email;type:varchar(100)" json:"email"`
	Password     string     `gorm:"column:password;type:varchar(200)" json:"-"`
	Semester     string     `gorm:"column:semester;type:enum('1','2','3','4','5','6','7','8')" json:"semester"`
	DepartmentId uint       `gorm:"column:department_id;type:uint" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentId" json:"-"`
	Address      string     `gorm:"column:address;default:null" json:"address,omitempty"`
	ProfileImage string     `gorm:"column:profile_image;default:null" json:"profile_image,omitempty"`
}

func (Student) TableName() string {
	return "students"
}
