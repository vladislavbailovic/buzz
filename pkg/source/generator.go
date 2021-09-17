package source

type Generator interface {
	HasNext() bool
	GetNext() string
	Size() int
	Reset()
}
