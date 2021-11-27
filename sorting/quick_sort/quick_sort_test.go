package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	actualResult := QuickSort([]int{10, -30, 0, 97, 5, -1})
	expectedResult := []int{-30, -1, 0, 5, 10, 97}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}
}
