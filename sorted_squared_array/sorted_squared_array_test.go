package sorted_squared_array

import "testing"

func TestSortedSquaredArray(t *testing.T) {
	expectedResult := []int{1, 4, 9, 25, 36, 64, 81}
	actualResult := SortedSquaredArray([]int{1, 2, 3, 5, 6, 8, 9})

	if len(actualResult) != len(expectedResult) {
		t.Fatalf("Expected %v but got %v", len(expectedResult), len(actualResult))
	}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}
}

func TestSortedSquaredArrayWithNegativeValues(t *testing.T) {
	expectedResult := []int{1, 4, 9, 16}
	actualResult := SortedSquaredArray([]int{-4, -3, -2, -1})

	if len(actualResult) != len(expectedResult) {
		t.Fatalf("Expected %v but got %v", len(expectedResult), len(actualResult))
	}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}
}
