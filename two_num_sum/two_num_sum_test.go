package two_num_sum

import "testing"

func TestTwoNumSum(t *testing.T) {
	expectedTotalSum := 10
	expectedResult := []int{-1, 11}
	actualResult := TwoNumSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)

	if len(actualResult) != len(expectedResult) {
		t.Fatalf("Expected %v but got %v", len(expectedResult), len(actualResult))
	}

	actualTotalSum := actualResult[0] + actualResult[1]
	if actualTotalSum != expectedTotalSum {
		t.Fatalf("Expected %v but got %v", expectedTotalSum, actualTotalSum)
	}
}
