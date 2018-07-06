/*
Package parsing turns the inputs into tokens
*/
package parsing

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jdholdren/airblog/pkg/token"
)

// Lexer is the struct to parse the file turn the file into its many tokens
type Lexer struct {
	r *bufio.Reader
}

// newLexer creates a lexer from the reader and emits tokens to the channel c
func newLexer(r io.Reader) *Lexer {
	return &Lexer{
		r: bufio.NewReader(r),
	}
}

func (l *Lexer) Run(init LexHandler, c chan token.Item) {
	f := init

	for {
		f = f(l, c)

		if f == nil {
			break
		}
	}

	// Clear up the channel
	close(c)
}

func (l *Lexer) Read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return rune(0)
	}

	return ch
}

// Peek reads the next rune, but then moves back
func (l *Lexer) Peek() rune {
	r, _, err := l.r.ReadRune()
	if err != nil {
		r = rune(0)
	}

	if r != rune(0) {
		l.Unread()
	}

	return r
}

func (l *Lexer) Unread() {
	if err := l.r.UnreadRune(); err != nil {
		fmt.Println(err)
	}
}

// LexHandler is the type for the recursive function
type LexHandler func(l *Lexer, c chan token.Item) LexHandler
