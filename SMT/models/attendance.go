package models

type Attendace struct {
	ID        uint    `gorm:"column:id;type:uint;primaryKey;autoIncrement" json:"id"`
	StudentID uint    `gorm:"column:student_id;type:uint;index:,unique,composite:Student_Date_Unique"`
	Student   Student `gorm:"foreignKey:StudentID" json:"-"`
	Date      string  `gorm:"column:date;type:date;index:,unique,composite:Student_Date_Unique" json:"date"`
	Status    bool    `gorm:"column:status" json:"status"`
}

func (Attendace) TableName() string {
	return "attendance"
}
