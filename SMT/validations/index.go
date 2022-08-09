package validations

import (
	stringTypes "SMT/types/strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidations() {
	v, _ := binding.Validator.Engine().(*validator.Validate)

	v.RegisterValidation(stringTypes.DEPARTMENT_CODE_VALIDATION, DepartmentCodeValidation)
	v.RegisterValidation(stringTypes.SEMESTER_VALIDATION, SemesterValidation)
	v.RegisterValidation(stringTypes.DATE_VALIDATION, _DateValidation)
}
