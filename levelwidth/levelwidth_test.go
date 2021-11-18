package levelwidth

import "testing"

func TestLevelWidthToReturnAnArray(t *testing.T) {
	n1 := NewNode(1)
	n1.Add(2)
	n1.Add(3)
	n1.Add(7)
	n1.Children[0].Add(4)
	n1.Children[0].Add(5)
	n1.Children[1].Add(6)
	n1.Children[0].Children[0].Add(8)

	actualResult := LevelWidth(n1)
	expectedResult := []int{1, 3, 3, 1}

	if len(actualResult) != 4 {
		t.Fatalf("Expected %v but got %v", 4, len(actualResult))
	}

	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", v, expectedResult[i])
		}
	}
}
