package lexer

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

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() Token {
	start := l.readPosition

	for isLetter(l.ch) {
		l.readChar()
	}

	return NewIdentifier(l.input[start:l.readPosition])
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readNumber() Token {
	start := l.readPosition

	for isDigit(l.ch) {
		l.readChar()
	}

	return NewInt(l.input[start:l.readPosition])
}

func (l *Lexer) NextToken() Token {
	var t Token
	l.skipWhitespace()

	if isLetter(l.ch) {
		t = l.readIdentifier()
	} else if isDigit(l.ch) {
		t = l.readNumber()
	} else {
		t = NewToken(l.ch)
		l.readChar() // Advance for the next read
	}

	return t
}
