package kata

func DirReduc(arr []string) []string {
	reduced := true
	for reduced {
		for i, _ := range arr {
			if isOpposite(arr[i], arr[i+1]) {
				reduced = true
				arr = append(arr[0:i], arr[i+1:]...)
			} else {
				reduced = false
			}
		}
	}
	return arr
}

func isOpposite(left, right string) bool {
	return left == "NORTH" && right == "SOUTH" ||
		left == "EAST" && right == "WEST" ||
		left == "SOUTH" && right == "NORTH" ||
		left == "WEST" && right == "EAST"
}
