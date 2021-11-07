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

func TestReverse1ToReturnReversedInputString(t *testing.T) {
	actualResults := []string{Reverse1("Hello"), Reverse1("apple"), Reverse1("Greetings!")}
	var expectedResults = []string{"olleH", "elppa", "!sgniteerG"}

	for i, v := range actualResults {

		if v != expectedResults[i] {
			t.Fatalf("Expected %s but got %s", expectedResults[i], v)
		}
	}
}

func TestReverse2ToReturnReversedInputString(t *testing.T) {
	actualResults := []string{Reverse2("Hello"), Reverse2("apple"), Reverse2("Greetings!")}
	var expectedResults = []string{"olleH", "elppa", "!sgniteerG"}

	for i, v := range actualResults {

		if v != expectedResults[i] {
			t.Fatalf("Expected %s but got %s", expectedResults[i], v)
		}
	}
}
