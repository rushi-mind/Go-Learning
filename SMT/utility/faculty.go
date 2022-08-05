package utility

import (
	"SMT/config"
	"SMT/models"
	"strconv"
)

func CreateEmployeeId(deptId uint) string {
	var lastTeacherId, departmentCode, employeeId string
	config.DB.Model(&models.Faculty{}).Select("employee_id").Order("id desc").Limit(1).First(&lastTeacherId)
	config.DB.Model(&models.Department{}).Select("code").Where("id = ?", deptId).Find(&departmentCode)
	if departmentCode == "" {
		return ""
	}
	employeeId = departmentCode
	if lastTeacherId == "" {
		return employeeId + "001"
	}
	temp, _ := strconv.Atoi(string([]byte(lastTeacherId)[len(departmentCode):]))
	temp++
	switch len(strconv.Itoa(temp)) {
	case 1:
		employeeId += "00" + strconv.Itoa(temp)
	case 2:
		employeeId += "0" + strconv.Itoa(temp)
	case 3:
		employeeId += strconv.Itoa(temp)
	}
	return employeeId
}

func CreateEmailForTeacher(empId string) string {
	return empId + "@smt.ac.com"
}
