package lexer

import "fmt"

type TokenType string

var (
	// operation tokens
	PLUS_OP  TokenType = "PLUS_OP"
	MINUS_OP TokenType = "MINUS_OP"
	MULT_OP  TokenType = "MULT_OP"
	DIV_OP   TokenType = "DIV_OP"
	MOD_OP   TokenType = "MOD_OP"
	POW_OP   TokenType = "POW_OP"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"

	NUMBER TokenType = "NUMBER"
)

// Token represents a token, and has a TokenType and a Value.
type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("%v: %s", t.Type, t.Value)
}
