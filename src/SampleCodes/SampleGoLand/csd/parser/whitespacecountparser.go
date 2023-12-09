package parser

import (
	"SampleGoLand/csd/parser/source"
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

}
