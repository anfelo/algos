package validate

import "testing"

func TestValidateBSTree(t *testing.T) {
	li := NewNode(-1, nil, nil)
	ri := NewNode(3, nil, nil)
	ni := NewNode(1, li, ri)
	ni.Insert(-2)
	ni.Insert(0)
	ni.Insert(10)
	ni.Insert(11)
	ni.Left.Right.Right = NewNode(2, nil, nil)

	if Validate(ni, nil, nil) {
		t.Fatalf("Expected %t but got %t", false, true)
	}

	lv := NewNode(-1, nil, nil)
	rv := NewNode(3, nil, nil)
	nv := NewNode(1, lv, rv)
	nv.Insert(-2)
	nv.Insert(0)
	nv.Insert(10)
	nv.Insert(11)

	if !Validate(nv, nil, nil) {
		t.Fatalf("Expected %t but got %t", true, false)
	}
}
