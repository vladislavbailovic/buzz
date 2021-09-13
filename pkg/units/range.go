package units

type Cursor struct {
	limit    int
	position int
}

func (r Cursor) IsEof() bool {
	return r.position >= r.limit
}
func (r *Cursor) Advance() {
	r.position += 1
}
func (r Cursor) GetPos() int {
	return r.position
}
func (r Cursor) GetSize() int {
	return r.limit
}
func NewCursor(size int) Cursor {
	return Cursor{size, 0}
}

type Bounds struct {
	min int
	max int
}

func (b Bounds) GetDiff() int {
	return b.max - b.min
}
func (b Bounds) GetMin() int {
	return b.min
}
func (b Bounds) GetMax() int {
	return b.max
}
func NewBounds(boundaries []int) Bounds {
	b := Bounds{}
	if len(boundaries) > 0 {
		b.max = boundaries[0]
	}
	if len(boundaries) > 1 {
		b.min = boundaries[1]
	}
	return b

}
