package kata

func ValidISBN10(isbn string) bool {
	if len(isbn)!= 10 {
		return false
	}
	var sum int
	for i := 0; i < 9; i++ {
		digit := int(isbn[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
        sum += digit * (i + 1)
    }
	if isbn[9] == 'X' || isbn[9] == 'x' {
		sum += 100
	} else {
		sum += int(isbn[9] - '0') * 10
	}
	return (sum % 11) == 0
}
