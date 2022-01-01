package valid_subsequence

import "testing"

func TestIsValidSubsequence(t *testing.T) {
	expectedResult := true
	actualResult := IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})

	if actualResult == false {
		t.Fatalf("Expected %t but got %t", expectedResult, actualResult)
	}
}

func TestIsValidSubsequenceNegative(t *testing.T) {
	expectedResult := false
	actualResult := IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8}, []int{1, 6, -1, 10})

	if actualResult == true {
		t.Fatalf("Expected %t but got %t", expectedResult, actualResult)
	}
}
