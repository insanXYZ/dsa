package sorting

func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}

		if !swapped {
			break
		}
	}
}
