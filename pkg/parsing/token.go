package parsing

import "fmt"

// ItemType is just an int, representing the different enumerations of token types
type itemType int

// The token constants
const (
	EOF    itemType = iota
	ItemH1          // H1 #
	ItemH2          // H2 ##
	ItemH3          // H3 ###
	L               // Literal String
)

// Item is the representation of the
type Item struct {
	// Type is the type of item, e.g. header, list item etc
	typ itemType

	// Value is the actual content of the token
	val string
}

// String lets the token be output
func (i Item) String() string {
	ending := ""
	if len(i.val) > 10 {
		ending = "..."
	}
	return fmt.Sprintf("%s: %.10s%s", i.typ.Name(), i.val, ending)
}

func (t itemType) Name() string {
	switch t {
	case ItemH1:
		return "h1"
	case ItemH2:
		return "h2"
	case ItemH3:
		return "h3"
	default:
		return "Unknown"
	}
}
