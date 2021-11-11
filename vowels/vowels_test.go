package vowels

import "testing"

func TestVowelsToReturnNumberOfVowelsInString(t *testing.T) {
	actualResults := []int{
		Vowels("Hi There!"),
		Vowels("Why do you asdk oh yeah?"),
		Vowels("Why?"),
	}
	var expectedResults = []int{3, 7, 0}

	for i, v := range actualResults {

		if v != expectedResults[i] {
			t.Fatalf("Expected %d but got %d", expectedResults[i], v)
		}
	}
}
