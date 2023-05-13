package chapter3

import "testing"

func Test_Point(t *testing.T) {
	a := NewPoint(1, 1)
	b := NewPoint(1, 1)
	if a != b {
		t.Fatalf("a and b were not equal")
	}

	// to human these two points are equal
}
