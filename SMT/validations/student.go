package validations

import (
	"github.com/go-playground/validator/v10"
)

var SemesterValidation validator.Func = func(fl validator.FieldLevel) bool {
	validSemesters := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	for _, val := range validSemesters {
		if fl.Field().String() == val {
			return true
		}
	}
	return false
}
