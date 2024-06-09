package heap

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestMaxHeapCorrectness(t *testing.T) {
	var (
		size  = int(rand.Int31n(1000) + 1000)
		items = make([]int, size)
		uut   = Make(func(a, b int) bool { return a < b }, size)
	)

	for i := range items {
		val := int(rand.Int31n(1000))
		items[i] = val
		uut.Push(val)
	}

	if uut.Len() != len(items) {
		t.Errorf("unexpected length: want=%d, got=%d", len(items), uut.Len())
	}

	sort.Slice(items, func(a, b int) bool { return items[a] > items[b] })

	for pos, want := range items {
		got, ok := uut.Pop()
		if !ok {
			break
		}
		if want != got {
			t.Errorf("pop returned unexpected value: pos=%d, want=%d, got=%d", pos, want, got)
		}
		if uut.Len() != len(items)-pos-1 {
			t.Errorf("unexpected length: want=%d, got=%d", 0, uut.Len())
		}
	}
}

func TestMinHeapCorrectness(t *testing.T) {
	var (
		size  = int(rand.Int31n(1000) + 1000)
		items = make([]int, size)
		uut   = Make(func(a, b int) bool { return a > b }, size) // less returns 'greater than' for min heap
	)

	for i := range items {
		val := int(rand.Int31n(1000))
		items[i] = val
		uut.Push(val)
	}

	if uut.Len() != len(items) {
		t.Errorf("unexpected length: want=%d, got=%d", len(items), uut.Len())
	}

	sort.Ints(items)

	for pos, want := range items {
		got, ok := uut.Pop()
		if !ok {
			break
		}
		if want != got {
			t.Errorf("pop returned unexpected value: pos=%d, want=%d, got=%d", pos, want, got)
		}
		if uut.Len() != len(items)-pos-1 {
			t.Errorf("unexpected length: want=%d, got=%d", 0, uut.Len())
		}
	}
}

func TestHeapStruct(t *testing.T) {
	type aStruct struct {
		compare   string
		dataInt   int
		dataFloat float64
	}

	var (
		size  = int(rand.Int31n(1000) + 1000)
		items = make([]aStruct, size)
		uut   = Make(func(a, b aStruct) bool { return a.compare < b.compare }, size)
	)

	for i := range items {
		r := int(rand.Int31n(1000))
		val := aStruct{
			compare:   strconv.Itoa(r),
			dataInt:   r,
			dataFloat: float64(r),
		}
		items[i] = val
		uut.Push(val)
	}

	if uut.Len() != len(items) {
		t.Errorf("unexpected length: want=%d, got=%d", len(items), uut.Len())
	}

	sort.Slice(items, func(a, b int) bool { return items[a].compare > items[b].compare })

	for pos, want := range items {
		got, ok := uut.Pop()
		if !ok {
			break
		}
		if want != got {
			t.Errorf("pop returned unexpected value: pos=%d, want=%v, got=%v", pos, want, got)
		}
		if uut.Len() != len(items)-pos-1 {
			t.Errorf("unexpected length: want=%d, got=%d", 0, uut.Len())
		}
	}
}

func TestHeapify(t *testing.T) {
	var (
		size = int(rand.Int31n(1000) + 1000)
		items = make([]int, size)
	)

	for i := range items {
		items[i] = int(rand.Int31n(1000))
	}

	uut, err := From(func(a, b int) bool { return a > b }, items...) // less returns 'greater than' for min heap
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if uut.Len() != len(items) {
		t.Errorf("unexpected length: want=%d, got=%d", len(items), uut.Len())
	}

	sort.Ints(items)

	for pos, want := range items {
		got, ok := uut.Pop()
		if !ok {
			break
		}
		if want != got {
			t.Errorf("pop returned unexpected value: pos=%d, want=%d, got=%d", pos, want, got)
		}
		if uut.Len() != len(items)-pos-1 {
			t.Errorf("unexpected length: want=%d, got=%d", 0, uut.Len())
		}
	}
}
