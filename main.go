package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shravanasati/odyssey/interpreter"
	"github.com/shravanasati/odyssey/lexer"
	"github.com/shravanasati/odyssey/parser"
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
		output, er := parser.Parse()
		// todo fix nil dereference on invalid syntax
		if er == nil && output != nil {
			fmt.Println(output.String())
			interpreter := interpreter.NewInterpreter()
			fmt.Println(interpreter.Visit(output))
		}
	}
}
