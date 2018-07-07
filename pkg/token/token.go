package token

import "fmt"

// ItemType is just an int, representing the different enumerations of token types
type ItemType int

// The token constants
const (
	EOF ItemType = iota
	H1           // H1 #
	H2           // H2 ##
	H3           // H3 ###
	BR           // A line break
	L            // Literal String
	I            // Italics
	B            // Bold
)

// Item is the representation of the
type Item struct {
	// Type is the type of item, e.g. header, list item etc
	Typ ItemType

	// Value is the actual content of the token
	Val string

	// Nested are the parts of this token that have been broken out
	Nested []Item
}

// String lets the token be output
func (i Item) String() string {
	ending := ""
	if len(i.Val) > 10 {
		ending = "..."
	}
	return fmt.Sprintf("%s: %.10s%s", i.Typ.Name(), i.Val, ending)
}

// Name gets the english version of the constant
func (t ItemType) Name() string {
	switch t {
	case H1:
		return "h1"
	case H2:
		return "h2"
	case H3:
		return "h3"
	case BR:
		return "break"
	case I:
		return "italics"
	case B:
		return "bold"
	case L:
		return "literal"
	default:
		return "Unknown"
	}
}
