package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Initialize with both currToken and peekToken set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}

	return prog
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	}

	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.nextTokenIs(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	if !p.nextTokenIs(token.ASSIGN) {
		return nil
	}

	// TODO: Handle the expression
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) nextTokenIs(tt token.TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	}

	return false
}

func (p *Parser) peekTokenIs(tt token.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) currTokenIs(tt token.TokenType) bool {
	return p.currToken.Type == tt
}
