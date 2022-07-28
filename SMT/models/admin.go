package models

type Admin struct {
	Id       uint   `gorm:"type:uint;primaryKey;column:id" json:"id"`
	AdminId  string `gorm:"type:varchar(12);column:adminId" json:"adminId"`
	EmailId  string `gorm:"type:varchar(30);column:emailId" json:"emailId"`
	Password string `gorm:"type:varchar(100);column:password" json:"password"`
}
