package lexer

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Lexer struct {
	input       string
	position    int    // current position in input (points to current char)
	currentChar string // current char under examination
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

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

func (l *Lexer) makeNumberToken() (Token, error) {
	var value string
	for isDigit(l.currentChar) || isDecimal(l.currentChar) {
		value += strings.TrimSpace(l.currentChar)
		if strings.Count(value, ".") > 1 {
			l.reportError(errMultipleDecimals)
			return Token{}, errMultipleDecimals
		}
		l.advance()
	}
	// add 0 at the start of a number if it starts with a decimal point
	// eg .52 => 0.52
	if string(value[0]) == "." {
		value = "0" + value
	}
	return Token{Type: NUMBER, Value: value}, nil
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

		case "e":
			tokens = append(tokens, Token{Type: NUMBER, Value: floatToString(math.E)})
			l.advance()

		case "p":
			if string(l.input[l.position + 1]) == "i" {
				tokens = append(tokens, 
					Token{Type: NUMBER, Value: floatToString(math.Pi)})
				l.advance()
				l.advance()
				// advancing twice because we've checked for two letters, pi
			} else {
				l.reportError(errInvalidCharacter)
				return []Token{}, errInvalidCharacter
			}

		default:
			if isDigit(l.currentChar) || isDecimal(l.currentChar) {
				numberToken, err := l.makeNumberToken()
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
