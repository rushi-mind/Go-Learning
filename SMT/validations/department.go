package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var DepartmentCodeValidation validator.Func = func(fl validator.FieldLevel) bool {
	regex, _ := regexp.Compile("^[0-9]{6,6}$")
	return regex.MatchString(fl.Field().String())
}
