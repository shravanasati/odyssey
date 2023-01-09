package lexer

import "errors"

// lexer errors defined here
var (
	errInvalidCharacter = errors.New("invalid character")
	errMultipleDecimals = errors.New("multiple decimals in a number")
)
