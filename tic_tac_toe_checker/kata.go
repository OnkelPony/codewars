package kata

const Rows = 3
const Cols = 3
const Points = 3

func IsSolved(board [3][3]int) int {
	var unifinished bool
  // check rows
	for row := 0; row < Rows; row++ {
		sum := 0
		for col := 0; col < Cols; col++ {
			sum += Pow(Points, board[row][col])
			if board[row][col] == 0 {
				unifinished = true
			}
		}
		if sum == 27 {
			return 2
		} else if sum == 9 {
			return 1
		}
	}
  // check columns
  for row := 0; row < Rows; row++ {
		sum := 0
		for col := 0; col < Cols; col++ {
			sum += Pow(Points, board[col][row])
		}
		if sum == 27 {
			return 2
		} else if sum == 9 {
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
