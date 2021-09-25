package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("odyssey> ")
		text, e := reader.ReadString('\n')
		if e != nil {
			break
		}
		fmt.Println(text)
	}
}