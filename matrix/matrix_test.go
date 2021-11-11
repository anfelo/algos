package matrix

import "testing"

func TestMatrixToReturnCorrectSequence1(t *testing.T) {
	actualResults := Matrix(2)
	var expectedResults = [][]int{
		{1, 2},
		{4, 3},
	}

	for i := range actualResults {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected level size %d but got %d", len(expectedResults[i]), len(actualResults[i]))
		}
		for j, v := range actualResults[i] {
			if v != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j], v)
			}
		}
	}
}

func TestMatrixToReturnCorrectSequence2(t *testing.T) {
	actualResults := Matrix(3)
	var expectedResults = [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}

	for i := range actualResults {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected level size %d but got %d", len(expectedResults[i]), len(actualResults[i]))
		}
		for j, v := range actualResults[i] {
			if v != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j], v)
			}
		}
	}
}

func TestMatrixToReturnCorrectSequence3(t *testing.T) {
	actualResults := Matrix(4)
	var expectedResults = [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	for i := range actualResults {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected level size %d but got %d", len(expectedResults[i]), len(actualResults[i]))
		}
		for j, v := range actualResults[i] {
			if v != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j], v)
			}
		}
	}
}
