package kata

func DirReduc(arr []string) []string {
	reduced := true
	for reduced {
		reduced = false
		for i := 0; i < len(arr) - 1; i++ {
			if isOpposite(arr[i], arr[i+1]) {
				reduced = true
				arr = append(arr[0:i], arr[i+2:]...)
				break
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
