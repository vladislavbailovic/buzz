package report

import (
	"buzz/pkg/challenge"
	"testing"
)

type FakeRejectFilter int

func (frf FakeRejectFilter) Apply(item Item) bool {
	return false
}

type FakeAcceptFilter int

func (faf FakeAcceptFilter) Apply(item Item) bool {
	return true
}

func TestAddResponseAddsItem(t *testing.T) {
	req := challenge.NewRequest("test.com")
	resp := challenge.Response{StatusCode: challenge.STATUS_PASSTHROUGH}
	report := Assembly{}

	if report.RawSize() != 0 {
		t.Fatalf("raw size should be zero initially")
	}

	report.Add(NewItem(req, resp))
	if report.RawSize() != 1 {
		t.Fatalf("raw size should be one after adding an item")
	}
}

func TestAddAddsItem(t *testing.T) {
	resp := challenge.Response{StatusCode: challenge.STATUS_PASSTHROUGH}
	req := challenge.NewRequest("test.com")
	item := NewItem(req, resp)
	report := Assembly{}

	if report.RawSize() != 0 {
		t.Fatalf("raw size should be zero initially")
	}

	report.Add(item)
	if report.RawSize() != 1 {
		t.Fatalf("raw size should be one after adding an item")
	}
}

func TestGetItemsAppliesFilters(t *testing.T) {
	resp := challenge.Response{StatusCode: challenge.STATUS_PASSTHROUGH}
	req := challenge.NewRequest("test.com")
	report := Assembly{}

	for i := 0; i < 10; i++ {
		report.Add(NewItem(req, resp))
	}

	if report.RawSize() != 10 {
		t.Fatalf("expected 10 items")
	}

	result1 := report.GetItems()
	if len(result1) != report.RawSize() {
		t.Fatalf("expected unfiltered result to be the same as raw size")
	}

	report.AddFilter(new(FakeAcceptFilter))
	result2 := report.GetItems()
	if len(result2) != report.RawSize() {
		t.Fatalf("passthrough filtered result should be the same as raw size")
	}

	report.AddFilter(new(FakeRejectFilter))
	result3 := report.GetItems()
	if len(result3) == report.RawSize() {
		t.Fatalf("filtered result should not be the same as raw size")
	}
}
