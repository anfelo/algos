package palindrome

import "testing"

func TestPalindromeToReturnTrueOrFalse(t *testing.T) {
	actualResults := []bool{Palindrome("abba"), Palindrome("apple")}
	var expectedResults = []bool{true, false}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %t but got %t", expectedResults[i], v)
		}
	}
}

func TestPalindrome1ToReturnTrueOrFalse(t *testing.T) {
	actualResults := []bool{Palindrome1("abba"), Palindrome1("apple")}
	var expectedResults = []bool{true, false}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %t but got %t", expectedResults[i], v)
		}
	}
}
