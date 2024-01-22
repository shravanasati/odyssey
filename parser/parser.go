package parser

import (
	"errors"
	"fmt"
	// "strings"

	"github.com/shravanasati/odyssey/lexer"
)

var errInvalidSyntax = errors.New("invalid syntax")

func itemInSlice[T comparable](item T, list []T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// todo write parser tests
type Parser struct {
	tokens       []lexer.Token
	position     int         // current position in tokens (points to current token)
	currentToken lexer.Token // current token under examination
}

func NewParser(tokens []lexer.Token) *Parser {
	parser := &Parser{
		tokens:   tokens,
		position: -1,
	}
	parser.advance()
	return parser
}

func (p *Parser) advance() {
	p.position++
	if p.position >= len(p.tokens) {
		p.currentToken = lexer.Token{}
	} else {
		p.currentToken = p.tokens[p.position]
	}
}

func (p *Parser) reportError(err error) {
	fmt.Printf("%s near %v \n", err.Error(), p.currentToken.String())
	// fmt.Printf(
	// 	"%s\n%s\n",
	// 	p.currentToken.String(),
	// 	strings.Repeat(" ", p.position)+"^ "+err.Error(),
	// )
	// fmt.Println(err.Error())
}

func (p *Parser) Parse() (Node, error) {
	if (p.currentToken == lexer.Token{}) {
		return nil, nil
	}
	result := p.expr()
	if (p.currentToken != lexer.Token{}) {
		p.reportError(errInvalidSyntax)
		return nil, errInvalidSyntax
	}
	return result, nil
}

func (p *Parser) expr() Node {
	result := p.term()
	operators := []lexer.TokenType{lexer.PLUS_OP, lexer.MINUS_OP}

	for (p.currentToken != lexer.Token{}) && itemInSlice(p.currentToken.Type, operators) {
		if p.currentToken.Type == lexer.PLUS_OP {
			p.advance()
			result = AddNode{result, p.term()}

		} else if p.currentToken.Type == lexer.MINUS_OP {
			p.advance()
			result = SubstractNode{result, p.term()}
		}
	}
	return result
}

func (p *Parser) term() Node {
	result := p.factor()
	operators := []lexer.TokenType{lexer.MOD_OP, lexer.EXP_OP, lexer.MULT_OP, lexer.DIV_OP}
	for (p.currentToken != lexer.Token{}) && itemInSlice(p.currentToken.Type, operators) {
		switch p.currentToken.Type {
		case lexer.MOD_OP:
			p.advance()
			result = ModulusNode{result, p.factor()}
		case lexer.EXP_OP:
			p.advance()
			result = ExponentNode{result, p.factor()}

		case lexer.MULT_OP:
			p.advance()
			result = MultiplyNode{result, p.factor()}
		case lexer.DIV_OP:
			p.advance()
			result = DivideNode{result, p.factor()}
		}
	}
	return result
}

func (p *Parser) factor() Node {
	token := p.currentToken
	switch token.Type {
	case lexer.LPAREN:
		p.advance()
		result := p.expr()
		if p.currentToken.Type != lexer.RPAREN {
			p.reportError(errInvalidSyntax)
			return nil
		}
		p.advance()
		return result
	case lexer.NUMBER:
		p.advance()
		return NumberNode{token.Value}
	case lexer.PLUS_OP:
		p.advance()
		return UnaryPlusNode{p.factor()}
	case lexer.MINUS_OP:
		p.advance()
		return UnaryMinusNode{p.factor()}
	default:
		p.reportError(errInvalidSyntax)
		return nil
	}
}
