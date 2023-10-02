package dyn_prog

func Fibo(n int, fibos map[int]int) int {
	if n == 0 || n == 1 {
	return n
	}
	if _, ok := fibos[n]; ok {
    return fibos[n]
    }
	fibos[n] = Fibo(n-1, fibos) + Fibo(n-2, fibos)
	return fibos[n]
}