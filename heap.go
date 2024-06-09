// Package heap is a generic heap implementation.
package heap

import "errors"

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// MakeHeap makes a heap. For a max heap, functional parameter 'less' must return true
// when item a < b and false otherwise. For a min heap, 'less' must return true when
// item a > b and false otherwise. 'cap' is the capcity guess of the maximum size of the
// heap and follows the same rules a capacity for a slice.
func Make[T any](less func(a, b T) bool, cap int) Heap[T] {
	return Heap[T]{
		data: make([]T, 0, cap),
		less: less,
	}
}

func From[T any](less func(a, b T) bool, elements ...T) (Heap[T], error) {
	var out Heap[T]

	if len(elements) == 0 {
		return out, nil
	}

	out.data = make([]T, len(elements), cap(elements))
	out.less = less

	if copy(out.data, elements) != len(elements) {
		return out, errors.New("copy slice failed")
	}

	for i := len(out.data)/2 - 1; i >= 0; i-- {
		out.heapify(i)
	}

	return out, nil
}

// i is the index of the root of a sub-tree to heapify
func (h *Heap[T]) heapify(i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < len(h.data) && !h.less(h.data[left], h.data[smallest]) {
		smallest = left
	}

	if right < len(h.data) && !h.less(h.data[right], h.data[smallest]) {
		smallest = right
	}

	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.heapify(smallest)
	}
}

// Cap returns the underlying capacity of the heap. This can grow.
func (h *Heap[T]) Cap() int {
	if h == nil {
		return 0
	}
	return cap(h.data)
}

// Len returns the number of items in the heap.
func (h *Heap[T]) Len() int {
	if h == nil {
		return 0
	}
	return len(h.data)
}

// Peek returns the top of the heap and true if there are items in the heap.
// It returns false when there are no items in the heap. Its runs in constant time.
func (h *Heap[T]) Peek() (T, bool) {
	if h != nil && len(h.data) > 0 {
		return h.data[0], true
	}
	var out T
	return out, false
}

// Pop returns the top item of the heap and removes this item from the heap. It
// returns true when there is at least 1 item in the heap and false otherwise.
// It runs in O(log n) where n is the size of the heap.
func (h *Heap[T]) Pop() (T, bool) {
	var out T
	if h == nil || len(h.data) == 0 {
		return out, false
	}

	out = h.data[0]
	if len(h.data) == 1 {
		h.data = h.data[:len(h.data)-1]
		return out, true
	}

	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	var (
		parent  = 0
		left    = 1
		right   = 2
		largest = parent
	)

	for left < len(h.data) {
		if h.less(h.data[parent], h.data[left]) {
			largest = left
		}
		if right < len(h.data) && h.less(h.data[largest], h.data[right]) {
			largest = right
		}
		if largest == parent {
			break
		}

		h.data[parent], h.data[largest] = h.data[largest], h.data[parent]

		parent = largest
		left = parent*2 + 1
		right = parent*2 + 2
	}

	return out, true
}

// Push adds a new item into the heap. It runs in time O(log n) where n is
// the size of the heap.
func (h *Heap[T]) Push(t T) {
	if h == nil || h.less == nil {
		return
	}

	h.data = append(h.data, t)

	cur := len(h.data) - 1
	parent := (cur - 1) / 2

	for parent != cur && !h.less(h.data[cur], h.data[parent]) {
		h.data[cur], h.data[parent] = h.data[parent], h.data[cur]

		cur = parent
		parent = (cur - 1) / 2
	}
}
