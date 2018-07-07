package markdown

import (
	"testing"
)

func TestBreakOutNoCases(t *testing.T) {
	// Test string
	lit := "This should have nothing special in it."

	res := breakOut(lit)

	// Should only have the one output
	if len(res) != 1 || res == nil {
		t.Fatal("Should have returned only one case")
	}

	i := res[0]

	// The first item should not have any items
	if len(i.Nested) > 0 {
		t.Fatal("Should not have any nested item")
	}

	// The value should be the same as the input
	if i.Val != lit {
		t.Fatal("Value was not the same as input")
	}
}
