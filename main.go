package main

import (
	"fmt"
	"io"
	"lang/lexer"
	"os"
)

func main() {
	file, err := os.Open("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	b, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	lxr := lexer.New()
	lxr.Parse(string(b))

	for next, hasNext := lxr.Iterator(); hasNext(); {
		token := next()
		fmt.Printf("[%d:%d]\t%s\t\t[%v]\n", token.Row, token.Column, token.Kind, token.Text)
	}
}
