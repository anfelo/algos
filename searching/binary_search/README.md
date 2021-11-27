# Binary Search

It only works if the array is already sorted. First start in the middle of the array, see if that is the value we are looking for if not go look in the upper/lower half of the array.

Runtime complexity: O(log(n))

# Pseudocode

* start at the middle of the array
* loop while array length is greater than 0
* if the middle is the value return true
* else if middle is less than value
	* discard the left half of the array
	* select a new middle from the right array
* return false

# Solution 1: (array slicing)

```go
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
```

# Solution 2: (min and max vars)

```go
func BinarySearch(n int, arr []int) bool {
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
```