package sorting

func QuickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		QuickSort(arr, low, p-1)
		QuickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[high], arr[i] = arr[i], arr[high]

	return i
}
