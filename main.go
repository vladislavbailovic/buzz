package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Range struct {
	limit    int
	position int
}

func (r Range) withinRange() bool {
	return r.position < r.limit
}
func (r *Range) advance() {
	r.position += 1
}

type Bounds struct {
	min int
	max int
}

func (b Bounds) getDiff() int {
	return b.max - b.min
}

type History struct {
	prev string
}

func (h *History) update(val string) {
	h.prev = val
}
func (h History) isKnown(val string) bool {
	return h.prev == val
}

type Generator interface {
	HasNext() bool
	GetNext() string
}

type RandomInt struct {
	rng     *Range
	bounds  Bounds
	history *History
}

func (ri RandomInt) HasNext() bool {
	return ri.rng.withinRange()
}
func (ri RandomInt) GetNext() string {
	var result int
	var val string
	rand.Seed(time.Now().UnixNano())
	delta := ri.bounds.getDiff()
	for val == "" || ri.history.isKnown(val) {
		if delta < 1 {
			result = rand.Int()
		} else {
			result = rand.Intn(delta) + ri.bounds.min
		}
		val = fmt.Sprintf("%d", result)
	}
	ri.history.update(val)
	ri.rng.advance()
	return val
}
