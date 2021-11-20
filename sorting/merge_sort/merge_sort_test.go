package mergesort

import "testing"

func TestMergeFunction(t *testing.T) {
	actualResult := Merge([]int{-25, 19, 50}, []int{0, 100})
	expectedResult := []int{-25, 0, 19, 50, 100}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}
}

func TestMergeSort(t *testing.T) {
	actualResult := MergeSort([]int{10, -30, 0, 97, 5, -1})
	expectedResult := []int{-30, -1, 0, 5, 10, 97}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}
}
