package source

type Source interface {
	NextCharacter() int
	Reset()
}
