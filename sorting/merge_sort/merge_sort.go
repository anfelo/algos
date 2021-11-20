package mergesort

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	h := len(arr) / 2
	a1, a2 := arr[0:h], arr[h:]

	return Merge(MergeSort(a1), MergeSort(a2))
}

func Merge(a1 []int, a2 []int) []int {
	var results []int
	for len(a1) > 0 && len(a2) > 0 {
		var l int
		if a1[0] < a2[0] {
			l, a1 = a1[0], a1[1:]
		} else {
			l, a2 = a2[0], a2[1:]
		}
		results = append(results, l)
	}
	results = append(results, a1...)
	results = append(results, a2...)
	return results
}
