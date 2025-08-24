package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := New(input)
	for i, et := range tests {
		nt := l.NextToken()

		if nt.Type != et.expectedType {
			t.Fatalf("test[%d] - incorrect tokentype. expected=%q, got=%q",
				i, et.expectedType, nt.Type)
		}

		if nt.Literal != et.expectedLiteral {
			t.Fatalf("test[%d] - incorrect literal. expected=%q, got=%q",
				i, et.expectedLiteral, nt.Literal)
		}
	}
}
