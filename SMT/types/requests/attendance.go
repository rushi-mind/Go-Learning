package requestTypes

type AttendanceInput struct {
	StudentID uint   `json:"student_id"`
	Date      string `json:"date" binding:"dateValidation"`
	Status    bool   `json:"status"`
}
