package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y
	};

	let result = add(five, ten);
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

		{LET, "let"},
		{IDENTIFIER, "add"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENTIFIER, "x"},
		{COMMA, ","},
		{IDENTIFIER, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{IDENTIFIER, "x"},
		{PLUS, "+"},
		{IDENTIFIER, "y"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},

		{LET, "let"},
		{IDENTIFIER, "result"},
		{ASSIGN, "="},
		{IDENTIFIER, "add"},
		{LPAREN, "("},
		{IDENTIFIER, "five"},
		{COMMA, ","},
		{IDENTIFIER, "ten"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{EOF, ""},
	}

	l := New(input)
	for i, et := range tests {
		nt := l.NextToken()

		// TODO: Consider factoring out the actual checking of types. This could let us
		//       have more granular tests.
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
