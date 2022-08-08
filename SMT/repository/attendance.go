package repository

import (
	"SMT/config"
	"SMT/models"
)

func InsertAttendance(requestBody []models.Attendace) error {
	return config.DB.Create(requestBody).Error
}
