package services

import "strings"

func CreateSlug(str string) string {
	return strings.ToLower(strings.Join(strings.Split(str, " "), "-"))
}
