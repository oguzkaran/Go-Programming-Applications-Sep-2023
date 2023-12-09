package parser

import "SampleGoLand/csd/parser/source"

type Parser interface {
	Parse(source source.Source)
}
