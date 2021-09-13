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

func NewStaticList(subject []string) StaticList {
	cursor := units.NewCursor(len(subject))
	return StaticList{
		&cursor,
		subject,
	}
}
