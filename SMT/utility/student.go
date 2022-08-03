package utility

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	"strconv"
	"strings"
	"time"
)

func CreateRollNo(student requestTypes.AddNewStudent, departmentCode string) string {
	rollNo := ""
	switch student.Semester {
	case "1", "2":
		rollNo += strconv.Itoa(time.Now().Year())
	case "3", "4":
		rollNo += strconv.Itoa(time.Now().Year() - 1)
	case "5", "6":
		rollNo += strconv.Itoa(time.Now().Year() - 2)
	case "7", "8":
		rollNo += strconv.Itoa(time.Now().Year() - 3)
	}
	rollNo += departmentCode
	var lastRollNo string
	config.DB.Model(&models.Student{}).Where("roll_no LIKE ?", rollNo+"%").Order("roll_no DESC").Limit(1).Select("roll_no").Find(&lastRollNo)
	if lastRollNo == "" {
		return rollNo + "001"
	}
	temp, _ := strconv.Atoi(strings.Split(lastRollNo, departmentCode)[1])
	switch len(strconv.Itoa(temp + 1)) {
	case 1:
		rollNo += "00" + strconv.Itoa(temp+1)
	case 2:
		rollNo += "0" + strconv.Itoa(temp+1)
	case 3:
		rollNo += strconv.Itoa(temp + 1)
	}
	return rollNo
}

func CreateEmailForStudent(rollNo string) string {
	return rollNo + "@smt.com"
}
