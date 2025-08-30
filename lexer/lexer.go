package lexer

import "monkey/token"

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

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()

	if isLetter(l.ch) {
		t = l.readIdentifier()
	} else if isDigit(l.ch) {
		t = l.readNumber()
	} else {
		t = l.readToken()
		l.readChar() // Advance for the next read
	}

	return t
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

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	}
	return l.input[l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.readPosition

	for isLetter(l.ch) {
		l.readChar()
	}

	return token.NewIdentifier(l.input[start:l.readPosition])
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readToken() token.Token {
	// Handle special, double character tokens.
	//
	// This makes the interface feel a bit awkward as token.go is no
	// longer responsible for taking in a value and spitting out a Token.
	//
	// The nice thing about this is that exceptions to token reading are
	// very explicit.
	if l.ch == '=' && l.peekChar() == '=' {
		l.readChar()
		return token.Token{Type: token.EQ, Literal: "=="}
	} else if l.ch == '!' && l.peekChar() == '=' {
		l.readChar()
		return token.Token{Type: token.NEQ, Literal: "!="}
	}

	return token.NewToken(l.ch)
}

func (l *Lexer) readNumber() token.Token {
	start := l.readPosition

	for isDigit(l.ch) {
		l.readChar()
	}

	return token.NewInt(l.input[start:l.readPosition])
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
