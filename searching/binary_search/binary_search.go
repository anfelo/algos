package binarysearch

func BinarySearch(n int, arr []int) bool {
	for len(arr) > 1 {
		h := len(arr) / 2
		m := arr[h]
		if m == n {
			return true
		} else if m > n {
			arr = arr[0:h]
		} else {
			arr = arr[h:]
		}
	}
	return false
}

func BinarySearch1(n int, arr []int) bool {
	min := 0
	max := len(arr) - 1
	for min <= max {
		h := (min + max) / 2
		m := arr[h]
		if m == n {
			return true
		} else if m > n {
			max = h - 1
		} else {
			min = h + 1
		}
	}
	return false
}
