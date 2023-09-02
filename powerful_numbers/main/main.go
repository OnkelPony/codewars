package kata
import "fmt"

func isDisarium(n int) bool {
  sum := 0
  for i, digit := range fmt.Sprint(n) {
    sum += int(digit) ^ (i + 1)
    fmt.Println("sum = ", sum)
  }
  return sum == n
}

func nextDisarium(n int) int {
  for i := n + 1; i < 1e6; i++ {
    if isDisarium(i) {
      return i
    }
  }
  return -1
}

func main() {
  fmt.Println(nextDisarium(2427))
}
