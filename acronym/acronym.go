package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	slice := regexp.MustCompile(`[^a-zA-Z]`).Split(s, -1)
	var result []string

	for _, value := range slice {
		charReg := regexp.MustCompile(`[a-zA-Z]`)
		result = append(result, strings.ToUpper(charReg.FindString(value)))
	}
	return strings.Join(result, "")
}
