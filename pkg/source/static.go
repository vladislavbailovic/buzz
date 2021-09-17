package source

import "buzz/pkg/units"

type StaticList struct {
	cursor *units.Cursor
	list   []string
}

func (lst StaticList) HasNext() bool {
	return !lst.cursor.IsEof()
}

func (lst StaticList) GetNext() string {
	val := lst.list[lst.cursor.GetPos()]
	lst.cursor.Advance()
	return val
}

func (lst StaticList) Reset() {
	lst.cursor.Reset()
}

func (lst StaticList) Size() int {
	return lst.cursor.GetSize()
}

func NewStaticList(subject []string) Generator {
	cursor := units.NewCursor(len(subject))
	return StaticList{
		&cursor,
		subject,
	}
}
