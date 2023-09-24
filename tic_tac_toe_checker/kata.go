package kata

const Rows = 3
const Cols = 3
const Points = 3

func IsSolved(board [3][3]int) int {
	var unifinished bool
  // check rows and columns
	for i := 0; i < Rows; i++ {
		sumR := 0
    sumC := 0
		for j := 0; j < Cols; j++ {
			sumR += Pow(Points, board[i][j])
      sumC += Pow(Points, board[j][i])
      if board[i][j] == 0 {
				unifinished = true
			}
		}
		if sumR == 27 || sumC == 27 {
			return 2
		} else if sumR == 9 || sumC == 9 {
			return 1
		}
	}
 
  // check diagonals
  sumB := 0
  sumS := 0
  for i := 0; i < Rows; i++ {
    sumB += Pow(Points, board[i][i])
    sumS += Pow(Points, board[i][Cols-1-i])
  }
  if sumB == 27 || sumS == 27 {
    return 2
  } else if sumB == 9 || sumS == 9 {
    return 1
  }
  if unifinished {
    return -1
  }
	return 0
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
