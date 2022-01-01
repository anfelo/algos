# Directions

Write a function that takes in a non-empty array of distinct integers and an
integer representing a target sum. If any two numbers in the input array sum
up to the target sum, the function should return them in an array, in any
order. If no two numbers sum up to the target sum, the function should return
an empty array.

Note that the target sum has to be obtained by summing two different integers
in the array; you can't add a single integer to itself in order to obtain the
target sum.

You can assume that there will be at most one pair of numbers summing up to
the target sum.

## Examples

- two_number_sum([3, 5, -4, 8, 11, 1, -1, 6], 10) == [-1, 11]

# Solution 1: double loop O(n^2)

- create a result array
- loop over the given array
	- save a reference to number in position i
	- loop over the elements in the array starting from i + 1
	- if element in position i plus element in position j equals the given sum
		- push the elements into the result array
		- return array
- return empty results array

```go
// O(n^2) time | O(n^2) space
func TwoNumSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}
```

# Solution 2: solving for y = total - x, O(n)

- create a hash table (a map of ints)
- loop over the elements in the array
	- solve y = total - currentNum
	- if y exists in the hash table (means that currentNum + y is equal to the total)
		- return an array with y and currentNum
	- else store currentNum in the hash table
- return an empty array

```go
// O(n) time | O(n) space
func TwoNumSum2(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}
```