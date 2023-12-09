package str

type StringSource struct {
	str   string
	index int
}

func New(str string) *StringSource {
	return &StringSource{str, 0}
}

func (ss *StringSource) NextCharacter() int {
	if ss.index == len(ss.str) {
		return -1
	}

	result := int(ss.str[ss.index])
	ss.index++
	return result
}

func (ss *StringSource) Reset() {
	ss.index = 0
}
