package valid_subsequence

// O(n) time, O(1) space
func IsValidSubsequence(array []int, sequence []int) bool {
	seqPointer := 0
	for i := 0; i < len(array); i++ {
		if array[i] == sequence[seqPointer] {
			seqPointer++
		}

		if seqPointer == len(sequence) {
			return true
		}
	}
	return false
}

func IsValidSubsequence2(array []int, sequence []int) bool {
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
