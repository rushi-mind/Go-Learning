package models

type Admin struct {
	Id       uint   `gorm:"type:uint;primaryKey;autoIncrement;column:id" json:"id"`
	AdminId  string `gorm:"type:varchar(12);column:admin_id" json:"admin_id"`
	EmailId  string `gorm:"type:varchar(30);column:email" json:"email"`
	Password string `gorm:"type:varchar(100);column:password" json:"password"`
}

func (Admin) TableName() string {
	return "admins"
}
