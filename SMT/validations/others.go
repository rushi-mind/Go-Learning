package validations

import (
	stringTypes "SMT/types/strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func DateValidation(date string) error {
	_, err := time.Parse(stringTypes.DATE_FORMAT, date)
	return err
}

var _DateValidation validator.Func = func(fl validator.FieldLevel) bool {
	_, err := time.Parse(stringTypes.DATE_FORMAT, fl.Field().String())
	return err == nil
}

func IsValidSemester(semester string) bool {
	return semester >= "1" && semester <= "8"
}
