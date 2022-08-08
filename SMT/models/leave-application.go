package models

import "time"

type LeaveApplication struct {
	ID         uint      `gorm:"column:id;type:uint;primaryKey;autoIncrement" json:"id"`
	StudentID  uint      `gorm:"column:student_id;type:uint" json:"student_id"`
	Student    Student   `gorm:"foreignKey:StudentID" json:"-"`
	DateFrom   time.Time `gorm:"column:date_from;type:date" json:"date_from"`
	DateTo     time.Time `gorm:"column:date_to;type:date" json:"date_to"`
	Duration   int       `gorm:"column:duration;type:int" json:"duration"`
	Reason     string    `gorm:"column:reason;type:text" json:"reason"`
	IsApproved bool      `gorm:"column:is_approved;type:boolean;default:false" json:"is_approved"`
}

func (LeaveApplication) TableName() string {
	return "leaveApplications"
}
