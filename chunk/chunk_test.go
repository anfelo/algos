package chunk

import "testing"

func TestChunkToReturnCorrectChucks(t *testing.T) {
	actualResults := Chunk([]int{1, 2, 3, 4}, 2)
	var expectedResults = [][]int{{1, 2}, {3, 4}}

	for i := 0; i < len(actualResults); i++ {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected chunk size %d but got %d", len(expectedResults[i]), len(actualResults[i]))
		}

		for j := 0; j < len(actualResults[i]); j++ {
			if actualResults[i][j] != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j], actualResults[i][j])
			}
		}
	}
}

func TestChunkToReturnCorrectChucks1(t *testing.T) {
	actualResults := Chunk([]int{1, 2, 3, 4, 5}, 2)
	var expectedResults = [][]int{{1, 2}, {3, 4}, {5}}

	for i := 0; i < len(actualResults); i++ {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected chunk size %d but got %d",
				len(expectedResults[i]), len(actualResults[i]))
		}

		for j := 0; j < len(actualResults[i]); j++ {
			if actualResults[i][j] != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j],
					actualResults[i][j])
			}
		}
	}
}

func TestChunkToReturnCorrectChucks2(t *testing.T) {
	actualResults := Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)
	var expectedResults = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}

	for i := 0; i < len(actualResults); i++ {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected chunk size %d but got %d",
				len(expectedResults[i]), len(actualResults[i]))
		}

		for j := 0; j < len(actualResults[i]); j++ {
			if actualResults[i][j] != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j],
					actualResults[i][j])
			}
		}
	}
}

func TestChunkToReturnCorrectChucks3(t *testing.T) {
	actualResults := Chunk([]int{1, 2, 3, 4, 5}, 4)
	var expectedResults = [][]int{{1, 2, 3, 4}, {5}}

	for i := 0; i < len(actualResults); i++ {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected chunk size %d but got %d",
				len(expectedResults[i]), len(actualResults[i]))
		}

		for j := 0; j < len(actualResults[i]); j++ {
			if actualResults[i][j] != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j],
					actualResults[i][j])
			}
		}
	}
}

func TestChunkToReturnCorrectChucks4(t *testing.T) {
	actualResults := Chunk([]int{1, 2, 3, 4, 5}, 10)
	var expectedResults = [][]int{{1, 2, 3, 4, 5}}

	for i := 0; i < len(actualResults); i++ {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected chunk size %d but got %d",
				len(expectedResults[i]), len(actualResults[i]))
		}

		for j := 0; j < len(actualResults[i]); j++ {
			if actualResults[i][j] != expectedResults[i][j] {
				t.Fatalf("Expected %d but got %d", expectedResults[i][j],
					actualResults[i][j])
			}
		}
	}
}
