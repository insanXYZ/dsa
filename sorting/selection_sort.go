package sorting

func SelectionSort(arr []int) {

	for i := range len(arr) - 1 {

		minIdx := i

		for j := 1 + i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
