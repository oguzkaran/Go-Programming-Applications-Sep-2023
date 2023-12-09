package app

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/parser"
	"SampleGoLand/csd/parser/source/str"
)

func RunWhitespaceCountParserApp() {
	wp := parser.New()

	for {
		s := console.ReadString("Input text:")
		if s == "quit" {
			break
		}

		ss := str.New(s)

		wp.Parse(ss)

		console.WriteLine("Count:%d", wp.Count())
	}
}
