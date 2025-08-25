package lexer

import (
	"testing"
)

type expectedToken struct {
	expectedType    TokenType
	expectedLiteral string
}

func assertTokenizes(t *testing.T, l *Lexer, tests []expectedToken) {
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

func TestNextTokenOperators(t *testing.T) {
	input := `
	!+-/*=
	5 < 10 > 5
	`

	tests := []expectedToken{
		{BANG, "!"},
		{PLUS, "+"},
		{MINUS, "-"},
		{SLASH, "/"},
		{ASTERISK, "*"},
		{ASSIGN, "="},

		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{GT, ">"},
		{INT, "5"},

		{EOF, ""},
	}

	l := New(input)
	assertTokenizes(t, l, tests)
}

func TestNextTokenVariable(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	`

	tests := []expectedToken{
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
	assertTokenizes(t, l, tests)
}

func TestNextTokenFunction(t *testing.T) {
	input := `
	let add = fn(x, y) {
		x + y
	};

	let result = add(5, 10);
	`

	tests := []expectedToken{
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
		{INT, "5"},
		{COMMA, ","},
		{INT, "10"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{EOF, ""},
	}

	l := New(input)
	assertTokenizes(t, l, tests)
}
