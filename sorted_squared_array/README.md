# Directions

Write a function that takes in a non-empty array of integers that are sorted in
ascending order and returns a new array of the same length with the squares of
the original integers also sorted in ascending order. (watch out for the
negative numbers!!)

## Examples

- sorted_squared_array([1, 2, 3, 5, 6, 8, 9]) == [1, 4, 9, 25, 36, 64, 81]

# Solution 1: loop + sort

## Pseudocode

- create an array of length equal to the input array
- loop over all the elements in the input array
  - store in the created array the squared value of item in position i
- sort in ascending order the items in the array
- return the sorted array

```go
// O(n*log(n)) time, O(n) space
func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))
	for i, v := range array {
		sortedSquares[i] = v * v
	}
	sort.Ints(sortedSquares)

	return sortedSquares
}
```

# Solution 2: loop & order while looping

## Pseudocode

- create an array of length equal to the input array
- create a variable that holds a reference to the index of the smallest value in
  the array
- create a variable that holds a reference to the index of the largest value in
  the array
- while the small index is less or equal to the largest index
  - calculate the squared value in the small index
  - calculate the squared value in the large index
  - if the small squared value is greater than the large squared value
    - add the small square to the beginning of the sorted array
    - add 1 to the small index
  - else
    - add the large square to the beginning of the sorted array
    - substract 1 from the large index
- return the sorted array

```go
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
```
