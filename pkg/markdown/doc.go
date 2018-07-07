/*
Package markdown represents the logic necessary to translate md files into tokens
*/
package markdown

import (
	"fmt"

	"github.com/jdholdren/airblog/pkg/parsing"
	"github.com/jdholdren/airblog/pkg/token"
)

// Initial is the starter function for markdown
func Initial(l *parsing.Lexer, c chan token.Item) parsing.LexHandler {
	switch l.Peek() {
	case '#':
		return parseHeader
	case rune(0):
		return nil
		// TODO Add cases for other types
	default:
		fmt.Println("Unable to handle")
		return nil
	}
}
