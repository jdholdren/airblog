package markdown

import (
	"github.com/jdholdren/airblog/pkg/parsing"
	"github.com/jdholdren/airblog/pkg/token"
)

func parseHeader(l *parsing.Lexer, c chan token.Item) parsing.LexHandler {
	// Get the number of '#' to determine the header type
	// Maxes out at h3
	count := 0
	for {
		r := l.Read()

		if r != '#' {
			l.Unread()
			break
		}

		count++
	}

	// Now ingore all whitespace until the end
	val := ""
	for {
		r := l.Read()

		if r == '\n' || r == rune(0) {
			break
		}

		if r != ' ' || val != "" {
			val += string(r)
		}
	}

	if val != "" {
		// Emit the element
		var typ token.ItemType
		switch count {
		case 1:
			typ = token.H1
		case 2:
			typ = token.H2
		default:
			typ = token.H3
		}

		c <- token.Item{
			Typ: typ,
			Val: val,
		}
	}

	return Initial
}
