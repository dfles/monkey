package lexer

import (
	"monkey/token"
	"testing"
)

type expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testTokenizes(t *testing.T, l *Lexer, tests []expectedToken) {
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

	10 == 10
	1 != 2
	`

	tests := []expectedToken{
		{token.BANG, "!"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.ASSIGN, "="},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},

		{token.INT, "1"},
		{token.NEQ, "!="},
		{token.INT, "2"},

		{token.EOF, ""},
	}

	l := New(input)
	testTokenizes(t, l, tests)
}

func TestNextTokenVariable(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	`

	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)
	testTokenizes(t, l, tests)
}

func TestNextTokenFunction(t *testing.T) {
	input := `
	let add = fn(x, y) {
		x + y
	};

	let result = add(5, 10);
	`

	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.COMMA, ","},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)
	testTokenizes(t, l, tests)
}

func TestNextTokenControlFlow(t *testing.T) {
	input := `
	let gt = fn(l, r) {
		if (l > r) {
			return true;
		}

		return false;
	};
	`

	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "gt"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "l"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "r"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "l"},
		{token.GT, ">"},
		{token.IDENTIFIER, "r"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)
	testTokenizes(t, l, tests)
}
