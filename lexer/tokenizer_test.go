package lexer

import (
	"math"
	"testing"
)

type lexerTest struct {
	input   string
	tokens  []Token
	wantErr bool
}

func Test_isDigit(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"1", true},
		{"0", true},
		{"0", true},
		{"regnerhiueb", false},
		{"", false},
		{"646464", true},
	}

	for _, test := range tests {
		if isDigit(test.input) != test.expected {
			t.Errorf("isDigit(%q) returned %t, expected %t", test.input, !test.expected, test.expected)
		}
	}
}

func Test_Tokenizer(t *testing.T) {
	tests := []lexerTest{
		{
			"", []Token{}, false,
		},

		{
			"   \t \t \n  \r",
			[]Token{}, false,
		},

		{
			"1 / 8 *     9", []Token{
				{NUMBER, "1"},
				{DIV_OP, "/"},
				{NUMBER, "8"},
				{MULT_OP, "*"},
				{NUMBER, "9"},
			}, false,
		},

		{
			"5+5", []Token{
				{NUMBER, "5"},
				{PLUS_OP, "+"},
				{NUMBER, "5"},
			}, false,
		},

		{
			"4545 ^ 19 % 27 + 121.211 - 55", []Token{
				{NUMBER, "4545"},
				{POW_OP, "^"},
				{NUMBER, "19"},
				{MOD_OP, "%"},
				{NUMBER, "27"},
				{PLUS_OP, "+"},
				{NUMBER, "121.211"},
				{MINUS_OP, "-"},
				{NUMBER, "55"},
			}, false,
		},

		{
			"(1 + 2) * 3 / ( 2 - 23)", []Token{
				{LPAREN, "("},
				{NUMBER, "1"},
				{PLUS_OP, "+"},
				{NUMBER, "2"},
				{RPAREN, ")"},
				{MULT_OP, "*"},
				{NUMBER, "3"},
				{DIV_OP, "/"},
				{LPAREN, "("},
				{NUMBER, "2"},
				{MINUS_OP, "-"},
				{NUMBER, "23"},
				{RPAREN, ")"},
			}, false,
		},

		{
			"(1 + (2 - 1)) * ( 5 / 7 ^ 8 % 4 )",
			[]Token{
				{LPAREN, "("},
				{NUMBER, "1"},
				{PLUS_OP, "+"},
				{LPAREN, "("},
				{NUMBER, "2"},
				{MINUS_OP, "-"},
				{NUMBER, "1"},
				{RPAREN, ")"},
				{RPAREN, ")"},
				{MULT_OP, "*"},
				{LPAREN, "("},
				{NUMBER, "5"},
				{DIV_OP, "/"},
				{NUMBER, "7"},
				{POW_OP, "^"},
				{NUMBER, "8"},
				{MOD_OP, "%"},
				{NUMBER, "4"},
				{RPAREN, ")"},
			}, false,
		},

		{
			"pi *e + .25", []Token{
				{NUMBER, floatToString(math.Pi)},
				{MULT_OP, "*"},
				{NUMBER, floatToString(math.E)},
				{PLUS_OP, "+"},
				{NUMBER, "0.25"},
			},
			false,
		},

		{
			"this is gonna fail", []Token{}, true,
		},

		{
			"12.23 + 89.254.1", nil, true,
		},
	}

	for _, test := range tests {
		l := NewLexer(test.input)

		result, err := l.Tokenize()
		if err != nil && !test.wantErr {
			t.Errorf("Tokenize(%q) returned error %q", test.input, err)
		}
		if err == nil && test.wantErr {
			t.Errorf("Tokenize(%q) did not return error", test.input)
		}

		if !test.wantErr {
			if len(result) != len(test.tokens) {
				t.Errorf("Tokenize(%q) returned %d tokens, expected %d", test.input, len(result), len(test.tokens))
			}
			for i, token := range result {
				if token.Type != test.tokens[i].Type {
					t.Errorf("Tokenize(%q) returned token %q, expected %q", test.input, token.Type, test.tokens[i].Type)
				}
				if token.Value != test.tokens[i].Value {
					t.Errorf("Tokenize(%q) returned token %q, expected %q", test.input, token.Value, test.tokens[i].Value)
				}
			}
		}
	}
}
