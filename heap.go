// Package heap is a generic max heap implementation.
package heap

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// MakeHeap makes a heap. For a max heap, functional parameter 'less' must return true
// when item a < b and false otherwise. For a min heap, 'less' must return true when 
// item a > b and false otherwise. 'cap' is the capcity guess of the maximum size of the
// heap.
func MakeHeap[T any](less func(a, b T) bool, cap int) Heap[T] {
	return Heap[T]{
		data: make([]T, 0, cap),
		less: less,
	}
}

// Cap returns the underlying capacity of the heap. This can grow.
func (h Heap[T]) Cap() int {
	return cap(h.data)
}

// Len returns the number of items in the heap.
func (h Heap[T]) Len() int {
	return len(h.data)
}

// Peek returns the top of the heap and true if there are items in the heap.
// It returns false when there are no items in the heap. Its runs in constant time.
func (h *Heap[T]) Peek() (T, bool) {
	if len(h.data) > 0 {
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
	if len(h.data) == 0 {
		return out, false
	}

	out = h.data[0]
	if len(h.data) == 1 {
		h.data = h.data[:len(h.data)-1]
		return out, true
	}

	h.data[0] = h.data[len(h.data) - 1]
	h.data = h.data[:len(h.data)-1]

	var (
		parent = 0
		left = 1
		right = 2
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
		left = parent * 2 + 1
		right = parent * 2 + 2
	}

	return out, true
}

// Push adds a new item into the heap. It runs in time O(log n) where n is
// the size of the heap.
func (h *Heap[T]) Push(t T) {
	h.data = append(h.data, t)

	cur := len(h.data) - 1
	parent := (cur - 1) / 2

	for parent != cur && !h.less(h.data[cur], h.data[parent]) {
		h.data[cur], h.data[parent] = h.data[parent], h.data[cur]

		cur = parent
		parent = (cur - 1) / 2
	}
}
