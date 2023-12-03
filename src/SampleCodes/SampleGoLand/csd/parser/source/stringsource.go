package source

type StringSource struct {
	str   string
	index int
}

func New(str string) *StringSource {
	return &StringSource{str, 0}
}

func (ss *StringSource) NextChar() int { //override
	return -1
}
