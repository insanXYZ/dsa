package searching

func LinearSearch(arr []int, x int) int {
	for i := range arr {
		if arr[i] == x {
			return i
		}
	}

	return -1
}
