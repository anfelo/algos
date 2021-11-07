package reverseint

import "testing"

func TestReverseIntToReturnReversedInputInt(t *testing.T) {
	actualResults := []int{Reverse(15), Reverse(981), Reverse(500), Reverse(-15), Reverse(-90)}
	var expectedResults = []int{51, 189, 5, -51, -9}

	for i, v := range actualResults {

		if v != expectedResults[i] {
			t.Fatalf("Expected %d but got %d", expectedResults[i], v)
		}
	}
}
