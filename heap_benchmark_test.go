package heap

import (
	stdheap "container/heap"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkIntPush(b *testing.B) {
	uut := Make(func(a, b int) bool { return a > b }, 4096)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := int(rand.Int31n(1000))
		uut.Push(val)
	}
}

// An IntHeap is a min heap of ints, copied from standard library docs.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkContainerIntPush(b *testing.B) {
	uut := make(IntHeap, 0, 4096)
	stdheap.Init(&uut)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := int(rand.Int31n(1000))
		uut.Push(val)
	}
}

func BenchmarkIntPop(b *testing.B) {
	uut := Make(func(a, b int) bool { return a > b }, 4096)
	for i := 0; i < b.N; i++ {
		val := int(rand.Int31n(1000))
		uut.Push(val)
	}

	total := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, ok := uut.Pop()
		if ok {
			total += val // ensure compiler doesn't optimize away the return value
		}
	}
}

func BenchmarkContainerIntPop(b *testing.B) {
	uut := make(IntHeap, 0, 4096)
	stdheap.Init(&uut)

	for i := 0; i < b.N; i++ {
		val := int(rand.Int31n(1000))
		stdheap.Push(&uut, val)
	}

	total := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := stdheap.Pop(&uut)
		if v, ok := val.(int); ok {
			total += v
		}
	}
}

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func BenchmarkStructPush(b *testing.B) {
	uut := Make(func(a, b *Item) bool { return a.priority > b.priority }, 4096)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := int(rand.Int31n(1000))

		val := Item{
			priority: r,
			value:    strconv.Itoa(r),
			index:    i,
		}

		uut.Push(&val)
	}
}

func BenchmarkContainerStructPush(b *testing.B) {
	uut := make(PriorityQueue, 0, 4096)
	stdheap.Init(&uut)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := int(rand.Int31n(1000))

		val := Item{
			priority: r,
			value:    strconv.Itoa(r),
			index:    i,
		}

		stdheap.Push(&uut, &val)
	}
}

func BenchmarkStructPop(b *testing.B) {
	uut := Make(func(a, b *Item) bool { return a.priority > b.priority }, 4096)

	for i := 0; i < b.N; i++ {
		r := int(rand.Int31n(1000))

		val := Item{
			priority: r,
			value:    strconv.Itoa(r),
			index:    i,
		}

		uut.Push(&val)
	}

	total := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, ok := uut.Pop()
		if ok {
			total += val.priority // ensure compiler doesn't optimize away the return value
		}
	}
}

func BenchmarkContainerStructPop(b *testing.B) {
	uut := make(PriorityQueue, 0, 4096)
	stdheap.Init(&uut)

	for i := 0; i < b.N; i++ {
		r := int(rand.Int31n(1000))

		val := Item{
			priority: r,
			value:    strconv.Itoa(r),
			index:    i,
		}

		stdheap.Push(&uut, &val)
	}

	total := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := stdheap.Pop(&uut)
		if v, ok := val.(int); ok {
			total += v
		}
	}
}
