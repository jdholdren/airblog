package markdown

import (
	"testing"

	"github.com/jdholdren/airblog/pkg/token"
)

func TestBreakOutNoCases(t *testing.T) {
	lit := "This should have nothing special in it."

	res := breakOut(lit)

	if len(res) != 1 || res == nil {
		t.Fatal("Should have returned only one case")
	}

	i := res[0]

	if len(i.Nested) > 0 {
		t.Fatal("Should not have any nested item")
	}

	if i.Val != lit {
		t.Fatal("Value was not the same as input")
	}
}

func TestBreakOutOneItalics(t *testing.T) {
	lit := "This should have *one* italics section."

	res := breakOut(lit)

	l := len(res)
	if l != 3 {
		t.Fatalf("Should contain 3 items, contains only %d", l)
	}

	if res[0].Typ != token.L {
		t.Fatalf("First should be type 'literal', was type %s", res[0].Typ.Name())
	}
	sh := "This should have "
	if res[0].Val != sh {
		t.Fatalf("First value should be \"%s\", was \"%s\"", sh, res[0].Val)
	}

	if res[1].Typ != token.I {
		t.Fatalf("Second should be type 'italics', was type %s", res[1].Typ.Name())
	}
	sh = "one"
	if res[1].Val != sh {
		t.Fatalf("Second value should be \"%s\", was \"%s\"", sh, res[1].Val)
	}

	if res[2].Typ != token.L {
		t.Fatalf("Third should be type 'literal', was type %s", res[2].Typ.Name())
	}
	sh = " italics section."
	if res[2].Val != sh {
		t.Fatalf("Third value should be \"%s\", was \"%s\"", sh, res[2].Val)
	}
}
