package lexer

import "testing"

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
			"1 / 8 * 9", []Token{
				{DIGIT, "1"},
				{DIV_OP, "/"},
				{DIGIT, "8"},
				{MULT_OP, "*"},
				{DIGIT, "9"},
			}, false,
		},

		{
			"4545 ^ 19 % 27 + 121211 - 55", []Token{
				{DIGIT, "4545"},
				{POW_OP, "^"},
				{DIGIT, "19"},
				{MOD_OP, "%"},
				{DIGIT, "27"},
				{PLUS_OP, "+"},
				{DIGIT, "121211"},
				{MINUS_OP, "-"},
				{DIGIT, "55"},
			}, false,
		},

		{
			"(1 + 2) * 3 / ( 2 - 23)", []Token{
				{LPAREN, "("},
				{DIGIT, "1"},
				{PLUS_OP, "+"},
				{DIGIT, "2"},
				{RPAREN, ")"},
				{MULT_OP, "*"},
				{DIGIT, "3"},
				{DIV_OP, "/"},
				{LPAREN, "("},
				{DIGIT, "2"},
				{MINUS_OP, "-"},
				{DIGIT, "23"},
				{RPAREN, ")"},
			}, false,
		},

		{
			"(1 + (2 - 1)) * ( 5 / 7 ^ 8 % 4 )",
			[]Token{
				{LPAREN, "("},
				{DIGIT, "1"},
				{LPAREN, "("},
				{DIGIT, "2"},
				{MINUS_OP, "-"},
				{DIGIT, "1"},
				{RPAREN, ")"},
				{RPAREN, ")"},
				{MULT_OP, "*"},
				{LPAREN, "("},
				{DIGIT, "5"},
				{DIV_OP, "/"},
				{DIGIT, "7"},
				{POW_OP, "^"},
				{DIGIT, "8"},
				{MOD_OP, "%"},
				{DIGIT, "4"},
				{RPAREN, ")"},
			}, false,
		},

		{
			"this is gonna fail", []Token{}, true,
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
