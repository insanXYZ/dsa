package searching

func BinarySearch(arr []int, low, high, x int) int {

	for low <= high {

		mid := low + (high-low)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return -1
}
