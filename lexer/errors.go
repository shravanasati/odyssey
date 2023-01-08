package lexer

import "errors"

// lexer errors defined here
var (
	invalidCharacter = errors.New("invalid character")
	multipleDecimals = errors.New("multiple decimals in a number")
)