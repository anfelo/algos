# BubbleSort

Find the greatest element in the array and move it to the far right side of the array

## Pseudocode

* from i=0 to i < array length
	* from j=0 to array length - i
		* if the elementat j is greater that j+1
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