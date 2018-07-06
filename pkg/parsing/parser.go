package parsing

import (
	"fmt"
	"io"

	"github.com/jdholdren/airblog/pkg/token"
)

// Parser is the struct for getting all logic from tokens into
type Parser struct {
	l *Lexer
}

// NewParser creates a new parser and makes all dependencies for the parser struct
func NewParser(r io.Reader) *Parser {
	return &Parser{
		l: newLexer(r),
	}
}

// Run turns the file into html :)
func (p *Parser) Run(init LexHandler) error {
	// Make the channel
	c := make(chan token.Item)

	// Tell the lexer to run
	go p.l.Run(init, c)

	for item := range c {
		fmt.Println(item)
	}

	return nil
}
