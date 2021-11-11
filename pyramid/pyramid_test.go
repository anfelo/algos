package pyramid

import "testing"

func TestPyramidToReturnTheCorrectSequence(t *testing.T) {
	actualResults := [][]string{
		Pyramid1(1),
		Pyramid1(2),
		Pyramid1(3),
		Pyramid1(4),
	}
	var expectedResults = [][]string{
		{"#"},
		{"-#-", "###"},
		{"--#--", "-###-", "#####"},
		{"---#---", "--###--", "-#####-", "#######"},
	}

	for i, v := range actualResults {
		if len(actualResults[i]) != len(expectedResults[i]) {
			t.Fatalf("Expected pyramid levels %d but got %d",
				len(expectedResults[i]), len(actualResults[i]))
		}
		for j, s := range v {
			if s != expectedResults[i][j] {
				t.Fatalf("Expected %s but got %s", expectedResults[i][j], s)
			}
		}
	}
}

func TestPyramidRecursiveToPrintTheCorrectSequence(t *testing.T) {
	Pyramid2(1, "", 0)
	Pyramid2(2, "", 0)
	Pyramid2(3, "", 0)
	Pyramid2(4, "", 0)
}
