# BubbleSort

Find the greatest element in the array and move it to the far right side of the array

Best case complexity: O(n) - sorted array
Average case complexity: O(n^2)
Worst case complexity: O(n^2)

## Pseudocode

* from i=0 to i < array length
	* from j=0 to array length - i
		* if the element at j is greater than j+1
			* swap elements at j and j+1

```go
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
```

# Solution 2: (swapped var)

```go
func BubbleSort(input []int) []int {
	// n is the number of items in our list
	n := len(input)
	// set swapped to true
	swapped := true
	// loop
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 0; i < n-1; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i] > input[i+1] {
				// swap values using Go's tuple assignment
				input[i], input[i+1] = input[i+1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}

	return input
}
```