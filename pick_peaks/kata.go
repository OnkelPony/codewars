package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	var result PosPeaks
	for i := 1; i < len(array)-1; i++ {
		j := i
		for ; j < len(array) - 1 && array[i] == array[j]; j++ {}
		if array[i] > array[i-1] && array[i] > array[j] {
			result.Pos = append(result.Pos, i)
			result.Peaks = append(result.Peaks, array[i])
		}
	}
	return result
}
