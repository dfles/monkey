package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	if prog == nil {
		t.Fatalf("ParseProgram() return nil")
	}

	checkParserErrors(t, p)

	if len(prog.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got=%d", len(prog.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral - expected: 'let', got: %q", s.TokenLiteral())
	}
	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("Statement - expected: '*ast.LetStatement', got: %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Statement.Name.Value - expected: '%s', got:'%s'", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Statement.Name.TokenLiteral - expected: %s, got: %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkParserErrors(t, p)

	if len(prog.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got=%d", len(prog.Statements))
	}

	for i, stmt := range prog.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Statement [%d] - expected: *ast.ReturnStatement, got: %T", i, returnStmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("TokenLiteral() [%d] - expected: 'return', got: %s", i, returnStmt.TokenLiteral())
		}

	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}
