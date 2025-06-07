package sorting

func merge(x, y []int) []int {
	var res []int

	i, j := 0, 0

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			res = append(res, x[i])
			i++
		} else {
			res = append(res, y[j])
			j++
		}
	}

	for ; i < len(x); i++ {
		res = append(res, x[i])
	}

	for ; j < len(y); j++ {
		res = append(res, y[j])
	}

	return res
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	l := MergeSort(arr[:len(arr)/2])
	r := MergeSort(arr[len(arr)/2:])

	return merge(l, r)
}
