package validations

import (
	"SMT/types"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidations() {
	v, _ := binding.Validator.Engine().(*validator.Validate)

	v.RegisterValidation(types.DEPARTMENT_CODE_VALIDATION, DepartmentCodeValidation)
	v.RegisterValidation(types.NAME_VALIDATION, NameValidation)
	v.RegisterValidation(types.SEMESTER_VALIDATION, SemesterValidation)
}
