package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	actualResult := BinarySearch(5, []int{-30, -1, 1, 3, 4, 5, 11})
	expectedResult := true

	if actualResult != expectedResult {
		t.Fatalf("Expected %t but got %t", expectedResult, actualResult)
	}
}

func TestNegativeBinarySearch(t *testing.T) {
	actualResult := BinarySearch(5, []int{-30, 1, 3, 4, 6, 11, 15})
	expectedResult := false

	if actualResult != expectedResult {
		t.Fatalf("Expected %t but got %t", expectedResult, actualResult)
	}
}
