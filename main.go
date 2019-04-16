package main

import (
	"log"

	"github.com/jtarchie/functional/parser"
)

//go:generate ragel -Z parser/parser.rl -o parser/parser.go
func main() {
	err := parser.Parse("age = 123")
	if err != nil {
		log.Fatalf("parser error: %s", err)
	}
}
