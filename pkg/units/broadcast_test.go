package units

import (
	"testing"
)

type FakeConsumer struct {
	data string
}

func (fc *FakeConsumer) Listen(what ...interface{}) {
	fc.data = what[0].(string)
}

func TestInit(t *testing.T) {
	cast := Broadcaster{}
	sub := new(FakeConsumer)
	cast.Subscribe("test", sub)

	if sub.data != "" {
		t.Fatalf("initial consumer data should be empty")
	}
	cast.Publish("test", "whatever")
	if sub.data != "whatever" {
		t.Fatalf("post-event consumer data should be set")
	}
}
