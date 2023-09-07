package kata

func Rot13(msg string) string {
	var r rune
	var result string
	for _, ch := range msg {
		if ch >= 'a' && ch <= 'm' || ch >= 'A' && ch <= 'M' {
			r = ch + 13
		} else if ch >= 'n' && ch <= 'z' || ch >= 'N' && ch <= 'Z' {
			r = ch - 13
		} else {
			r = ch
		}
		result += string(r)
	}
	return result
}