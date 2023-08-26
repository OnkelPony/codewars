package kata

func Cats(start, finish int) int {
	steps := (finish - start) / 3
	steps += (finish - start) % 3
	return steps
}
