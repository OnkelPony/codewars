package kata

import (
	"strconv"
	"strings"
)

func RGB(r, g, b int) string {
	var result string
	for _, color := range []int{r, g, b} {
		switch {
		case color <= 0:
			color = 0
		case color > 255:
			color = 255
		default:
		}
		result += strings.ToUpper(strconv.FormatInt(int64(color), 16))
	}
	return result
}
