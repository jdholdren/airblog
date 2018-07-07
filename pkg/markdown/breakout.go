package markdown

import (
	"github.com/jdholdren/airblog/pkg/token"
)

// Parses a string into smaller tokens, breaking up high level tokens into
// 'bold', 'italics', etc.
func breakOut(s string) []token.Item {
	// Make the array to return
	items := []token.Item{}

	if s == "" {
		return items
	}

	items = append(items, token.Item{
		Typ: token.L,
		Val: s,
	})

	return items
}
