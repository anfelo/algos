package insertionsort

import "testing"

func TestInsertionSort(t *testing.T) {
	actualResult := InsertionSort([]int{10, -30, 0, 97, 5})
	expectedResult := []int{-30, 0, 5, 10, 97}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}
}
