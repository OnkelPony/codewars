package kata

func FindNb(m int) int {
	var wannabeM int
	n := 0
	for wannabeM < m {
		n++
		wannabeM += n * n * n
	}
	if wannabeM == m {
		return n
	}
	return -1
}
