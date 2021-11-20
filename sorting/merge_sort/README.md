# MergeSort

MergeSort is solved commonly using recursion. It uses a helper function 'merge' that takes 2 sorted arrays and merge them into 1 sorted array

## Pseudocode

### Merger: Merge Function

* create results array
* while there are still elements in both arrays
	* if first element in the left half is less than first in right half
		* 'shift' the element from left into a 'result' arr
	* else 'shift' the element from the right into a 'result' arr
* take everything from the array that still has stuff in it and put it in results

```go
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
```

### MergeSort

* if the array has length 1
	* return array
* split the array into 2
* call mergesort with a1 and assign value to a1
* call mergesort with a2 and assign value to a2
* call the merge function with a1 and a2

```go
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	h := len(arr) / 2
	a1, a2 := arr[0:h], arr[h:]

	return Merge(MergeSort(a1), MergeSort(a2))
}
```