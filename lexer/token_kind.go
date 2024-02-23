package lexer

type TokenKind int

func (k TokenKind) String() string {
	return TokenKinds[k]
}

func GetTokenKind(char string) TokenKind {
	switch char {
	case ".":
		return T_DOT
	case ",":
		return T_COMMA

	case "[":
		return T_OPEN_SQUIRE_BRACKET
	case "]":
		return T_CLOSE_SQUIRE_BRACKET
	case "{":
		return T_OPEN_BRACE
	case "}":
		return T_CLOSE_BRACE
	case "(":
		return T_OPEN_BRACKET
	case ")":
		return T_CLOSE_BRACKET

	case "\"":
		return T_QUOTE
	case "'":
		return T_SINGLE_QUOTE

	case "=":
		return T_EQUIALS
	case "+":
		return T_PLUS
	case "-":
		return T_MINUS
	case "/":
		return T_SLASH
	case "%":
		return T_PERCENT
	case "*":
		return T_ASTERISK

	case "&":
		return T_AMPERSAND
	case "~":
		return T_TILDE
	case "<":
		return T_LESS_THAN
	case ">":
		return T_GREATER_THAN

	case "?":
		return T_QUESTION
	case "!":
		return T_EXCLAMATION

	case " ":
		return T_SPACE

	case "\r":
		return T_CARRIAGE_RETURN
	case "\n":
		return T_LINE_FEED

	default:
		return unknown
	}
}
