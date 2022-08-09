package repository

import (
	"SMT/config"
	requestTypes "SMT/types/requests"
	"strconv"
)

func InsertAttendance(requestBody []requestTypes.AttendanceInput) error {
	sql := "INSERT INTO attendance(student_id, date, status) VALUES "
	for i, obj := range requestBody {
		sql += `(` + strconv.Itoa(int(obj.StudentID)) + `, "` + obj.Date + `", ` + strconv.FormatBool(obj.Status) + `)`
		if i < len(requestBody)-1 {
			sql += ", "
		}
	}
	sql += ";"
	return config.DB.Exec(sql).Error
}
