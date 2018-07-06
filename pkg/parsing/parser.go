package parsing

import (
	"fmt"
	"io"
)

// Parser is the struct for getting all logic from tokens into
type Parser struct {
	l *lexer
}

// NewParser creates a new parser and makes all dependencies for the parser struct
func NewParser(r io.Reader) *Parser {
	return &Parser{
		l: newLexer(r),
	}
}

// Run turns the file into html :)
func (p *Parser) Run() error {
	// Make the channel
	c := make(chan Item)

	// Tell the lexer to run
	go p.l.run(c)

	for item := range c {
		fmt.Println(item)
	}

	return nil
}
