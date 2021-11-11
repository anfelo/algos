package steps

import "testing"

func TestStepsToReturnTheCorrectSequence(t *testing.T) {
	actualResults := [][]string{
		Steps1(2),
		Steps1(3),
		Steps1(4),
	}
	var expectedResults = [][]string{
		{"# ", "##"},
		{"#  ", "## ", "###"},
		{"#   ", "##  ", "### ", "####"},
	}

	for i, v := range actualResults {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected step size %d but got %d",
				len(expectedResults[i]), len(actualResults[i]))
		}
		for j, s := range v {
			if s != expectedResults[i][j] {
				t.Fatalf("Expected %s but got %s", expectedResults[i][j], s)
			}
		}
	}
}

func TestStepsRecursiveToPrintTheCorrectSequence(t *testing.T) {
	Steps2(2, 0)
	Steps2(3, 0)
	Steps2(4, 0)
}
