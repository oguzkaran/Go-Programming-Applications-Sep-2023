package app

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/parser"
	"SampleGoLand/csd/parser/source/str"
	"fmt"
)

func RunWhitespaceCountParserApp() {
	fmt.Println("Whitespace counter")
	wp := parser.New()

	for {
		s := console.ReadString("Input text:")
		if s == "quit" {
			break
		}

		ss := str.New(s)

		wp.Parse(ss)

		fmt.Printf("Number of whitespaces:%d\n", wp.Count())
	}
}
