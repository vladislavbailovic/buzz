package source

import (
	"fmt"
	"math/rand"
	"time"

	"buzz/pkg/units"
)

func NewRandomInt(size int, boundaries ...int) RandomInt {
	cursor := units.NewCursor(size)
	history := units.NewHistory()
	return RandomInt{
		&cursor,
		units.NewBounds(boundaries),
		&history,
	}
}

type RandomInt struct {
	cursor  *units.Cursor
	bounds  units.Bounds
	history *units.History
}

func (ri RandomInt) HasNext() bool {
	return !ri.cursor.IsEof()
}
func (ri RandomInt) GetNext() string {
	var result int
	var val string
	rand.Seed(time.Now().UnixNano())
	delta := ri.bounds.GetDiff()
	for val == "" || ri.history.IsKnown(val) {
		if delta < 1 {
			result = rand.Int()
		} else {
			result = rand.Intn(delta) + ri.bounds.GetMin()
		}
		val = fmt.Sprintf("%d", result)
	}
	ri.history.Update(val)
	ri.cursor.Advance()
	return val
}
