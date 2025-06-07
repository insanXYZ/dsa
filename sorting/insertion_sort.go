package sorting

func InsertionSort(arr []int) {

	for i := range len(arr) {

		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key

	}

}
