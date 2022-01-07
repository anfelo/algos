package sorted_squared_array

import "sort"

// O(n*log(n)) time, O(n) space
func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))
	for i, v := range array {
		sortedSquares[i] = v * v
	}
	sort.Ints(sortedSquares)

	return sortedSquares
}

// O(n) time, O(n) space
func SortedSquaredArray2(array []int) []int {
	var sorted []int
	sm := 0
	l := len(array) - 1
	for sm <= l {
		smSquare := array[sm] * array[sm]
		lSquare := array[l] * array[l]
		if smSquare > lSquare {
			sorted = append([]int{smSquare}, sorted...)
			sm++
		} else {
			sorted = append([]int{lSquare}, sorted...)
			l--
		}
	}
	return sorted
}
