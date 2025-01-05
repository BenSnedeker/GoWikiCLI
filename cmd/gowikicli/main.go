package main

import (
	"fmt"
	"gowikicli/internal/parser"
	"os"
)

func main() {
	err := parser.HandleInput(os.Args)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
