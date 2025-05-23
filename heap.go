package ds

import "container/heap"

type Heap[T any] interface {
	Push(k T)
	Pop() T
	Top() T
	IsEmpty() bool
	Size() int
}

type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

func NewHeap[T any](t HeapType, less func(a, b T) bool, elements ...T) Heap[T] {
	l := &list[T]{
		elements: elements,
		less:     less,
		heapType: t,
	}
	h := &myHeap[T]{l: l}
	// Initialize the heap
	heap.Init(h.l)
	return h
}

type myHeap[T any] struct {
	l *list[T]
}

// IsEmpty implements Heap.
func (h *myHeap[T]) IsEmpty() bool {
	return h.l.Len() == 0
}

// Pop implements Heap.
func (h *myHeap[T]) Pop() T {
	if h.IsEmpty() {
		panic("no objects in heap")
	}
	ele := heap.Pop(h.l)
	return ele.(T)
}

// Push implements Heap.
func (h *myHeap[T]) Push(k T) {
	heap.Push(h.l, k)
}

// Size implements Heap.
func (h *myHeap[T]) Size() int {
	return h.l.Len()
}

// Top implements Heap.
func (h *myHeap[T]) Top() T {
	if h.IsEmpty() {
		panic("no objects in heap")
	}
	return h.l.elements[0]
}

type list[T any] struct {
	elements []T
	heapType HeapType
	less     func(a, b T) bool
}

// Pop implements heap.Interface.
func (l *list[T]) Pop() any {
	if len(l.elements) == 0 {
		panic("no objects in heap")
	}
	n := len(l.elements) - 1
	x := l.elements[n]
	l.elements = l.elements[:n]
	return x
}

// Push implements heap.Interface.
func (l *list[T]) Push(x any) {
	if x == nil {
		panic("cannot push nil")
	}
	l.elements = append(l.elements, x.(T))
}

func (l *list[T]) Len() int {
	return len(l.elements)
}

func (l *list[T]) Less(i, j int) bool {
	if l.heapType == MinHeap {
		return l.less(l.elements[i], l.elements[j])
	}
	return l.less(l.elements[j], l.elements[i])
}

func (l *list[T]) Swap(i, j int) {
	l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
}
