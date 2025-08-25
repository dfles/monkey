package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Meta
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"
	LT       = "LT"
	GT       = "GT"

	// Delimiters
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var tokens = map[byte]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	'-': MINUS,
	'!': BANG,
	'*': ASTERISK,
	'/': SLASH,
	'<': LT,
	'>': GT,

	',': COMMA,
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,

	0: EOF,
}

func NewToken(ch byte) Token {
	tt, ok := tokens[ch]

	if ok && tt == EOF {
		return Token{Type: EOF}
	} else if ok {
		return Token{Type: tt, Literal: string(ch)}
	}

	return Token{Type: ILLEGAL, Literal: string(ch)}
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func NewIdentifier(s string) Token {
	if tt, ok := keywords[s]; ok {
		return Token{Type: tt, Literal: s}
	}

	return Token{Type: IDENTIFIER, Literal: s}
}

func NewInt(s string) Token {
	return Token{Type: INT, Literal: s}
}
