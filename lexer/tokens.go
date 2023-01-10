package lexer

type TokenType string

var (
	// operation tokens
	PLUS_OP  TokenType 
	MINUS_OP TokenType 
	MULT_OP  TokenType 
	DIV_OP   TokenType 
	MOD_OP   TokenType 
	POW_OP   TokenType 

	LPAREN TokenType 
	RPAREN TokenType 

	NUMBER TokenType 

)
