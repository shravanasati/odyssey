package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/Shravan-1908/odyssey/lexer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("odyssey> ")

		text, e := reader.ReadString('\n')
		if e != nil {
			break
		}

		lexer := lexer.NewLexer(text)
		tokens, err := lexer.Tokenize()

		if err != nil {
			continue
		}

		for _, token := range tokens {
			fmt.Printf("%v: %v\n", token.Type, token.Value)
		}
	}
}