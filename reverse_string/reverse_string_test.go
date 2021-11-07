package reversestring

import "testing"

func TestReverseToReturnReversedInputString(t *testing.T) {
	actualResults := []string{Reverse("Hello"), Reverse("apple"), Reverse("Greetings!")}
	var expectedResults = []string{"olleH", "elppa", "!sgniteerG"}

	for i, v := range actualResults {

		if v != expectedResults[i] {
			t.Fatalf("Expected %s but got %s", expectedResults[i], v)
		}
	}
}
