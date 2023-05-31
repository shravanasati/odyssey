package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shravan-1908/odyssey/lexer"
	"github.com/Shravan-1908/odyssey/parser"
)

func main() {
	fmt.Println("Type '/q' to exit.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("odyssey> ")

		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(text)) == "/q" {
			break
		}

		lexer := lexer.NewLexer(text)
		tokens, err := lexer.Tokenize()

		if err != nil {
			continue
		}

		parser := parser.NewParser(tokens)
		output := parser.Parse()
		if output != nil {
			fmt.Println(output.String())
		}
	}
}
