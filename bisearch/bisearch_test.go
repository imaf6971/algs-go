package bisearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	item := 2
	want := 1
	idx := BinarySearch(list, item)
	if idx != want {
		t.Fatalf("Searching %v in %v, wanted: %v but got %v", item, list, want, idx)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	list := []int{1, 2, 3}
	item := 42
	want := -1
	idx := BinarySearch(list, item)
	if idx != want {
		t.Fatalf("Searching %v in %v, wanted: %v but got %v", item, list, want, idx)
	}
}
