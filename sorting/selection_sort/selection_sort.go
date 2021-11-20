package selectionsort

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		indexOfMin := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[indexOfMin] {
				indexOfMin = j
			}
		}
		if i != indexOfMin {
			arr[i], arr[indexOfMin] = arr[indexOfMin], arr[i]
		}
	}
	return arr
}
