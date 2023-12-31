package kata

import "strings"

func ValidBraces(str string) bool {
  	openingBraces := "([{"

  if len(str) == 0{
    return true
  }
  if len(str) %2 == 1 || !strings.Contains(openingBraces, string(str[0])){
	return false
  }
	bracePairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var braces []rune
	for _, ch := range str {
		if strings.Contains(openingBraces, string(ch)) {
			braces = append(braces, ch)
		} else if bracePairs[ch] == braces[len(braces)-1] {
			braces = braces[:len(braces)-1]
		} else {
			return false
		}
	}
	return len(braces) == 0
}
