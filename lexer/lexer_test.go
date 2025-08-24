package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LET, "let"},
		{IDENTIFIER, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},

		{LET, "let"},
		{IDENTIFIER, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
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
