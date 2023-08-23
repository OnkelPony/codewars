package kata

import "strings"

func ValidBraces(str string) bool {
	bracePairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	openingBraces := "([{"
	var braces []rune
	for _, ch := range str {
		if strings.Contains(openingBraces, string(ch)) {
			braces = append(braces, ch)
		} else if ch == bracePairs[braces.pop()] {
		}
	}
	return false
}
