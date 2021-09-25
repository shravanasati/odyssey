package lexer

import (
	"errors"
	"strconv"
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

func isDigit(c string) bool {
	_, err := strconv.ParseInt(c, 10, 64)
	if err != nil {
		_, err2 := strconv.ParseFloat(c, 64)
		return err2 == nil
	}
	return err == nil
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

		case "==":
			tokens = append(tokens, Token{Type: EQ_OP, Value: l.currentChar})
			l.advance()
		case "!=":
			tokens = append(tokens, Token{Type: NEQ_OP, Value: l.currentChar})
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
				token := Token{Type: DIGIT, Value: l.currentChar}
				for isDigit(l.currentChar) {
					l.advance()
					token.Value += l.currentChar
				}
				tokens = append(tokens, token)
				l.advance()
			} else {
				return []Token{}, errors.New("invalid character")
			}
		}
	}

	l.advance()
	return tokens, nil
}
