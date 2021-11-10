package fizzbuzz

import "testing"

func TestFizzbuzzToReturnCorrectSequence(t *testing.T) {
	actualResults := Fizzbuzz(15)
	var expectedResults = []string{
		"1",
		"2",
		"fizz",
		"4",
		"buzz",
		"fizz",
		"7",
		"8",
		"fizz",
		"buzz",
		"11",
		"fizz",
		"13",
		"14",
		"fizzbuzz",
	}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %s but got %s", expectedResults[i], v)
		}
	}
}
