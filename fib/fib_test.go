package fib

import "testing"

func TestFibToReturnTheNthEntry(t *testing.T) {
	actualResults := []int{
		Fib(4),
		Fib(8),
	}
	var expectedResults = []int{
		3,
		21,
	}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %d but got %d", expectedResults[i], v)
		}
	}
}

func TestFibRecToReturnTheNthEntry(t *testing.T) {
	actualResults := []int{
		FibRec(4),
		FibRec(8),
	}
	var expectedResults = []int{
		3,
		21,
	}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %d but got %d", expectedResults[i], v)
		}
	}
}

func TestMemFibToReturnTheNthEntry(t *testing.T) {
	fastFib := Mem(FibRec)
	actualResults := []int{
		fastFib(4),
		fastFib(8),
	}
	var expectedResults = []int{
		3,
		21,
	}

	for i, v := range actualResults {
		if v != expectedResults[i] {
			t.Fatalf("Expected %d but got %d", expectedResults[i], v)
		}
	}
}
