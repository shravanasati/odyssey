package lexer

import (
	"errors"
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
		strings.Repeat(" ", l.position) + "^ " + err.Error() + "\n",
	)
}

func isDigit(c string) bool {
	_, err := strconv.ParseFloat(c, 64)
	return err == nil
}

func (l *Lexer) makeDigitToken() Token {
	var value string
	for isDigit(l.currentChar) {
		value += strings.TrimSpace(l.currentChar)
		l.advance()
	}

	return Token{Type: DIGIT, Value: value}
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
				tokens = append(tokens, l.makeDigitToken())
				l.advance()
			} else {
				err := errors.New("invalid character")
				l.reportError(err)
				return []Token{}, err
			}
		}
	}

	l.advance()
	return tokens, nil
}
