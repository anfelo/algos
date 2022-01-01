# Directions

Given two non-empty arrays of integers, write a function that determines whether
the second array is a subsequence of the first one.

A subsequence of an array is a set of numbers that aren't necessarily adjacent
in the array but that are in the same order as they appear in the array. For
instance, the numbers [1, 3, 4] form a subsequence of the array [1, 2, 3, 4],
and so do the numbers [2, 4]. Note that a single number in an array and the
array itself are both valid subsequences of the array.

## Examples

- is_valid_subsequence([5, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1, 10]) == true

# Solution 1:

## Pseudocode

- initialize the a pointer to the index of the sequence
- loop over the array while not at the end of the sequence and the array
  - if the current value of the array is equal to the sequence in the index
    - move the sequence index by 1
  - move the array index by 1
- return true if reached the end of the sequence else false

```go
func IsValidSubsequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return seqIdx == len(sequence)
}
```
