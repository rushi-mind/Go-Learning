package utility

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(err error) string {
	e := err.(validator.ValidationErrors)
	if strings.Contains(e[0].ActualTag(), "required") {
		return fmt.Sprintf("%s is required", e[0].StructField())
	}
	return fmt.Sprintf("%s: '%v' is not valid", e[0].StructField(), e[0].Value())
}
