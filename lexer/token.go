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
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func NewToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}
