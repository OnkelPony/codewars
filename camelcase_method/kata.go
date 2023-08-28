package kata

import "strings"


func CamelCase(s string) string {
	s = strings.Title(s)
	return strings.ReplaceAll(s, " ", "")
}