package parser

import (
	"SampleGoLand/csd/parser/source"
	"unicode"
)

type WhitespaceCountParser struct {
	count int
}

func New() *WhitespaceCountParser {
	return &WhitespaceCountParser{}
}

func (wp *WhitespaceCountParser) Count() int {
	return wp.count
}

func (wp *WhitespaceCountParser) Parse(source source.Source) {
	source.Reset()

	for {
		c := source.NextCharacter()

		if c == -1 {
			return
		}

		if unicode.IsSpace(rune(c)) {
			wp.count++
		}
	}
}
