package kata

// LCS returns the longest common subsequence of a and b.
func LCS(s1, s2 string) string {

	L := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		L[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				L[i+1][j+1] = L[i][j] + 1
			} else if L[i+1][j] > L[i][j+1] {
				L[i+1][j+1] = L[i+1][j]
			} else {
				L[i+1][j+1] = L[i][j+1]
			}
		}
	}

	result := make([]byte, 0, L[len(s1)][len(s2)])
	for x, y := len(s1), len(s2); x != 0 && y != 0; {
		if L[x][y] == L[x-1][y] {
			x--
		} else if L[x][y] == L[x][y-1] {
			y--
		} else {
			result = append(result, s1[x-1])
			x--
			y--
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return  string(result)
}