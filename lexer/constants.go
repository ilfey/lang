package lexer

const (
	T_PLAIN TokenKind = iota
	T_NUMBER

	T_DOT   // "."
	T_COMMA // ","

	T_OPEN_SQUIRE_BRACKET  // "["
	T_CLOSE_SQUIRE_BRACKET // "]"
	T_OPEN_BRACE           // "{"
	T_CLOSE_BRACE          // "}"
	T_OPEN_BRACKET         // "("
	T_CLOSE_BRACKET        // ")"

	T_QUOTE        // "\""
	T_SINGLE_QUOTE // "'"

	T_EQUIALS  // "="
	T_PLUS     // "+"
	T_MINUS    // "-"
	T_SLASH    // "/"
	T_PERCENT  // "%"
	T_ASTERISK // "*"

	T_AMPERSAND    // "&"
	T_TILDE        // "~"
	T_LESS_THAN    // "<"
	T_GREATER_THAN // ">"

	T_QUESTION    // "?"
	T_EXCLAMATION // "!"

	T_SPACE // " "

	T_CARRIAGE_RETURN // "\r" temp => T_NEXT_LINE
	T_LINE_FEED       // "\n" temp => T_NEXT_LINE

	T_NEXT_LINE // "\n"

	unknown // temp => T_NUMBER or T_PLAIN
)

var TokenKinds = map[TokenKind]string{
	T_PLAIN:  "T_PLAIN",
	T_NUMBER: "T_NUMBER",

	T_DOT:   "T_DOT",
	T_COMMA: "T_COMMA",

	T_OPEN_SQUIRE_BRACKET:  "T_OPEN_SQUIRE_BRACKET",
	T_CLOSE_SQUIRE_BRACKET: "T_CLOSE_SQUIRE_BRACKET",
	T_OPEN_BRACE:           "T_OPEN_BRACE",
	T_CLOSE_BRACE:          "T_CLOSE_BRACE",
	T_OPEN_BRACKET:         "T_OPEN_BRACKET",
	T_CLOSE_BRACKET:        "T_CLOSE_BRACKET",

	T_QUOTE:        "T_QUOTE",
	T_SINGLE_QUOTE: "T_SINGLE_QUOTE",

	T_EQUIALS:  "T_EQUIALS",
	T_PLUS:     "T_PLUS",
	T_MINUS:    "T_MINUS",
	T_SLASH:    "T_SLASH",
	T_PERCENT:  "T_PERCENT",
	T_ASTERISK: "T_ASTERISK",

	T_AMPERSAND:    "T_AMPERSAND",
	T_TILDE:        "T_TILDE",
	T_LESS_THAN:    "T_LESS_THAN",
	T_GREATER_THAN: "T_GREATER_THAN",

	T_QUESTION:    "T_QUESTION",
	T_EXCLAMATION: "T_EXCLAMATION",

	T_SPACE: "T_SPACE",

	T_CARRIAGE_RETURN: "T_CARRIAGE_RETURN",
	T_LINE_FEED:       "T_LINE_FEED",

	T_NEXT_LINE: "T_NEXT_LINE",
}

var numbers = []string{
	"1", "2", "3", "4", "5", "6", "7", "9", "0",
}
