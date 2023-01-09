package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

type Lexer struct {
	input       string
	position    int    // current position in input (points to current char)
	currentChar string // current char under examination
}

type Token struct {
	Type  TokenType
	Value string
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: -1,
	}
}

func (l *Lexer) advance() {
	l.position++
	if l.position >= len(l.input) {
		l.currentChar = ""
	} else {
		l.currentChar = string(l.input[l.position])
	}
}

func (l *Lexer) reportError(err error) {
	fmt.Printf("%s%s",
		l.input,
		strings.Repeat(" ", l.position)+"^ "+err.Error()+"\n",
	)
}

func isDigit(c string) bool {
	_, err := strconv.ParseFloat(c, 64)
	return err == nil
}

func isDecimal(c string) bool {
	return c == "."
}

func (l *Lexer) makeDigitToken() (Token, error) {
	var value string
	for isDigit(l.currentChar) || isDecimal(l.currentChar) {
		value += strings.TrimSpace(l.currentChar)
		if strings.Count(value, ".") > 1 {
			l.reportError(errMultipleDecimals)
			return Token{}, errMultipleDecimals
		}
		l.advance()
	}

	return Token{Type: DIGIT, Value: value}, nil
}

func (l *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token

	for l.position < len(l.input) {
		switch l.currentChar {
		case "+":
			tokens = append(tokens, Token{Type: PLUS_OP, Value: l.currentChar})
			l.advance()
		case "-":
			tokens = append(tokens, Token{Type: MINUS_OP, Value: l.currentChar})
			l.advance()
		case "*":
			tokens = append(tokens, Token{Type: MULT_OP, Value: l.currentChar})
			l.advance()
		case "/":
			tokens = append(tokens, Token{Type: DIV_OP, Value: l.currentChar})
			l.advance()

		case "^":
			tokens = append(tokens, Token{Type: POW_OP, Value: l.currentChar})
			l.advance()
		case "%":
			tokens = append(tokens, Token{Type: MOD_OP, Value: l.currentChar})
			l.advance()

		case "(":
			tokens = append(tokens, Token{Type: LPAREN, Value: l.currentChar})
			l.advance()
		case ")":
			tokens = append(tokens, Token{Type: RPAREN, Value: l.currentChar})
			l.advance()

		case " ", "\n", "\t", "\r", "":
			l.advance()

		default:
			if isDigit(l.currentChar) {
				numberToken, err := l.makeDigitToken()
				if err != nil {
					return []Token{}, err
				}
				tokens = append(tokens, numberToken)
				// not advancing here because the makeDigitToken method
				// does it itself
				// l.advance()
			} else {
				l.reportError(errInvalidCharacter)
				return []Token{}, errInvalidCharacter
			}
		}
	}

	l.advance()
	return tokens, nil
}
