/*
Package parsing turns the inputs into tokens
*/
package parsing

import (
	"bufio"
	"io"
)

// Lexer is the struct to parse the file turn the file into its many tokens
type lexer struct {
	r *bufio.Reader
}

// newLexer creates a lexer from the reader and emits tokens to the channel c
func newLexer(r io.Reader) *lexer {
	return &lexer{
		r: bufio.NewReader(r),
	}
}

func (l *lexer) run(c chan Item) {
	for {
		r := l.read()

		c <- Item{
			typ: ItemH1,
			val: string(r),
		}

		if r == rune(0) {
			break
		}
	}

	// Clear up the channel
	close(c)
}

func (l *lexer) read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return rune(0)
	}

	return ch
}
