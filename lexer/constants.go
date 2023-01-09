package lexer

type TokenType string

const (
	// operation tokens
	PLUS_OP  TokenType = "PLUS_OP"
	MINUS_OP TokenType = "MINUS_OP"
	MULT_OP  TokenType = "MULT_OP"
	DIV_OP   TokenType = "DIV_OP"
	MOD_OP   TokenType = "MOD_OP"
	POW_OP   TokenType = "POW_OP"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"

	DIGIT TokenType = "DIGIT"
)
