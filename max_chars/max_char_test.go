package maxchar

import "testing"

func TestMaxCharToReturnMostCommonChar(t *testing.T) {
	actualResults := []string{MaxChar("abbaa"), MaxChar("abcdefg 12311")}
	var expectedResults = []string{"a", "1"}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %s but got %s", expectedResults[i], v)
		}
	}
}
