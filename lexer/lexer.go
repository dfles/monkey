package lexer

type Lexer struct {
	input    string
	position int
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() Token {
	var ch byte

	if l.position >= len(l.input) {
		ch = '0'
	} else {
		ch = l.input[l.position]
		l.position = l.position + 1
	}

	var t Token
	switch ch {
	case '=':
		t = NewToken(ASSIGN, ch)
	case '+':
		t = NewToken(PLUS, ch)
	case ',':
		t = NewToken(COMMA, ch)
	case ';':
		t = NewToken(SEMICOLON, ch)
	case '(':
		t = NewToken(LPAREN, ch)
	case ')':
		t = NewToken(RPAREN, ch)
	case '{':
		t = NewToken(LBRACE, ch)
	case '}':
		t = NewToken(RBRACE, ch)
	case '0':
		t = Token{Type: EOF}
	}

	return t
}
