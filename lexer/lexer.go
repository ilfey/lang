package lexer

type Position struct {
	Column int
	Row    int
	Index  int
}

type Token struct {
	*Position
	Kind TokenKind
	Text string
}

type Lexer struct {
	tokens []*Token
	source string
	*Position
}

func New() *Lexer {
	return &Lexer{
		Position: &Position{},
	}
}

func (l *Lexer) getChar() string {
	return string(l.source[l.Index])
}

func (l *Lexer) getPosition() *Position {
	return &Position{
		Column: l.Column,
		Row:    l.Row,
		Index:  l.Index,
	}
}

func (l *Lexer) Iterator() (next func() *Token, hasNext func() bool) {
	l.Index = -1

	hasNext = func() bool {
		return l.Index < len(l.tokens)-1
	}

	next = func() *Token {
		l.Index++

		return l.tokens[l.Index]
	}

	return
}

func (l *Lexer) ParseIterator(source string) (next func() *Token, hasNext func() bool) {
	l.source = source
	l.Index = -1
	l.Column = 0
	l.Row = 1

	hasNext = func() bool {
		return l.Index < len(l.source)-1
	}

	next = func() *Token {
		l.Index++
		l.Column++
		char := l.getChar()
		pos := l.getPosition()
		kind := GetTokenKind(char)

		switch kind {
		case T_CARRIAGE_RETURN:
			if l.parseNextLine() {
				l.Row++
				l.Column = 0
			}
			return nil
		case T_LINE_FEED:
			if l.parseNextLine() {
				l.Row++
				l.Column = 0
			}
			return nil
		default:
			if kind == unknown {
				if !l.parseNumbers() {
					l.Index--
				}

				if !l.parsePlain() {
					l.Index--
				}

				return nil
			}
		}

		return &Token{
			Position: pos,
			Kind:     kind,
			Text:     char,
		}
	}

	return
}

func (l *Lexer) Parse(source string) {
	for next, hasNext := l.ParseIterator(source); hasNext(); {
		token := next()

		if token == nil {
			continue
		}

		l.tokens = append(l.tokens, token)
	}
}

func (l *Lexer) parseNumbers() bool {
	token := &Token{
		Position: l.getPosition(),
		Kind:     T_NUMBER,
	}

	for ; l.Index < len(l.source); l.Index++ {
		char := l.getChar()

		if l.isNumber(char) {
			l.Column++
			token.Text += char
			continue
		}

		if len(token.Text) == 0 {
			return true
		}

		l.tokens = append(l.tokens, token)

		return true
	}

	l.tokens = append(l.tokens, token)

	return true
}

func (l *Lexer) isNumber(s string) bool {
	for _, n := range numbers {
		if n == s {
			return true
		}
	}

	return false
}

func (l *Lexer) parsePlain() bool {
	token := &Token{
		Position: l.getPosition(),
		Kind:     T_PLAIN,
	}

	for ; l.Index < len(l.source); l.Index++ {
		char := l.getChar()

		if GetTokenKind(char) != unknown {
			if len(token.Text) == 0 {
				return false
			}

			l.tokens = append(l.tokens, token)

			return false
		}

		l.Column++
		token.Text += char
	}

	l.tokens = append(l.tokens, token)

	return true
}

func (l *Lexer) parseNextLine() bool {
	token := &Token{
		Position: l.getPosition(),
		Kind:     T_NEXT_LINE,
	}

	for ; l.Index < len(l.source); l.Index++ {
		char := l.getChar()
		kind := GetTokenKind(char)

		if kind == T_CARRIAGE_RETURN {
			token.Text += "\\r"
			continue
		}

		if kind == T_LINE_FEED {
			token.Text += "\\n"
			break
		}
	}

	l.tokens = append(l.tokens, token)

	return true
}
