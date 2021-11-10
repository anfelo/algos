package capitalize

import "testing"

func TestCapitalizeToReturnCapitalizedString(t *testing.T) {
	actualResults := []string{
		Capitalize("hi there, how is it going?"),
		Capitalize("i love breakfast at bill miller bbq"),
	}
	var expectedResults = []string{
		"Hi There, How Is It Going?",
		"I Love Breakfast At Bill Miller Bbq",
	}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %s but got %s", expectedResults[i], v)
		}
	}
}
