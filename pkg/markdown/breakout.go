package markdown

import (
	"github.com/jdholdren/airblog/pkg/token"
)

const (
	italR = '*'
)

// Parses a string into smaller tokens, breaking up high level tokens into
// 'bold', 'italics', etc.
func breakOut(s string) []token.Item {
	// Make the array to return
	items := []token.Item{}

	if s == "" {
		return items
	}

	lastPos := 0
	italicsPos := -1

	for pos, r := range s {
		switch r {
		case italR: // Italics
			if italicsPos == -1 {
				// Move the position
				italicsPos = pos
			} else {
				// Split the string from here to the last pos on the italicsPos into two items
				left := token.Item{
					Typ: token.L,
					Val: s[lastPos:italicsPos],
				}
				right := token.Item{
					Typ: token.I,
					Val: s[italicsPos+1 : pos],
				}
				items = append(items, left, right)

				// Reset italicsPos
				italicsPos = -1

				// Reset lastPos
				lastPos = pos + 1
			}
		}
	}

	// If there's any leftover, add it as a final item
	if len(s)-1 > lastPos {
		items = append(items, token.Item{
			Typ: token.L,
			Val: s[lastPos:],
		})
	}

	return items
}
