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
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

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
)
