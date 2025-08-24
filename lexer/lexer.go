package lexer

// Could these Token constructors live in token?
func newToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

type Lexer struct {
	input        string
	position     int
	readPosition int // 0 or position - 1; used to read multi-char identifiers
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Initialize with first char read

	return l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}

	l.readPosition = l.position
	l.position = l.position + 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func (l *Lexer) readIdentifier() Token {
	start := l.readPosition

	for isLetter(l.ch) {
		l.readChar()
	}

	s := l.input[start:l.readPosition]
	if tt, ok := keywords[s]; ok {
		return Token{Type: tt, Literal: s}
	}

	return Token{Type: IDENTIFIER, Literal: s}
}

func (l *Lexer) readNumber() Token {
	start := l.readPosition

	for isDigit(l.ch) {
		l.readChar()
	}

	s := l.input[start:l.readPosition]
	return Token{Type: INT, Literal: s}
}

func (l *Lexer) NextToken() Token {
	var t Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		t = newToken(ASSIGN, l.ch)
	case '+':
		t = newToken(PLUS, l.ch)
	case ',':
		t = newToken(COMMA, l.ch)
	case ';':
		t = newToken(SEMICOLON, l.ch)
	case '(':
		t = newToken(LPAREN, l.ch)
	case ')':
		t = newToken(RPAREN, l.ch)
	case '{':
		t = newToken(LBRACE, l.ch)
	case '}':
		t = newToken(RBRACE, l.ch)
	case 0:
		t = Token{Type: EOF}
	default:
		if isLetter(l.ch) {
			return l.readIdentifier()
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			t = Token{Type: ILLEGAL}
		}
	}

	l.readChar()
	return t
}
