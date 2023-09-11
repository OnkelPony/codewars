package kata

import (
	"bytes"
	"fmt"
)

func RGB(r, g, b int) string {
	var buf bytes.Buffer
	for _, color := range []int{r, g, b} {
		switch {
		case color <= 0:
			color = 0
		case color > 255:
			color = 255
		default:
		}
		buf.WriteString(fmt.Sprintf("%02X", color))
	}
	return buf.String()
}
