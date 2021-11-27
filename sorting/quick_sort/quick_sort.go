package quicksort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	p := arr[len(arr)-1]
	var l []int
	var r []int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < p {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}

	result := QuickSort(l)
	result = append(result, p)
	result = append(result, QuickSort(r)...)
	return result
}
