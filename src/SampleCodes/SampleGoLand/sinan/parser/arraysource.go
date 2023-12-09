package parser

type ArraySource struct {
	//...
}

func New() *ArraySource {
	return &ArraySource{}
}

func (as *ArraySource) NextCharacter() int {
	return -1
}

func (as *ArraySource) Reset() {
	//...
}
