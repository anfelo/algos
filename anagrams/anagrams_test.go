package anagrams

import "testing"

func TestAnagramsToReturnTrueOrFalse(t *testing.T) {
	actualResults := []bool{
		Anagrams("rail safety", "fairy tales"),
		Anagrams("RAIL! SAFETY!", "fairy tales"),
		Anagrams("Hi there!", "Bye there"),
	}
	var expectedResults = []bool{true, true, false}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %v but got %v", expectedResults[i], v)
		}
	}
}
