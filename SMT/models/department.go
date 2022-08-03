package models

type Department struct {
	Id       uint      `gorm:"primaryKey;autoIncrement;column:id;type:uint" json:"id"`
	Code     string    `gorm:"type:varchar(6);column:code" json:"code"`
	Name     string    `gorm:"type:varchar(30);column:name" json:"name"`
	Slug     string    `gorm:"type:varchar(30);column:slug" json:"-"`
	Students []Student `gorm:"DepartmentId" json:"-"`
}
