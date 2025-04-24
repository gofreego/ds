package ds

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type linkedListIterator[T any] struct {
	current *Node[T]
}

func newLinkedListIterator[T any](head *Node[T]) Iterator[*Node[T]] {
	return &linkedListIterator[T]{current: head}
}

func (it *linkedListIterator[T]) Next() *Node[T] {
	if it.current == nil {
		panic("no more elements")
	}

	node := it.current
	it.current = it.current.Next
	return node
}

func (it *linkedListIterator[T]) HasNext() bool {
	return it.current != nil
}

type LinkedList[T any] interface {
	Head() *Node[T]
	Tail() *Node[T]
	Append(k T)
	Prepend(k T)
	Remove(node *Node[T])
	Size() int
	IsEmpty() bool
	Iterator() Iterator[*Node[T]]
}

type linkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Append implements LinkedList.
func (l *linkedList[T]) Append(k T) {
	if l.head == nil {
		l.head = &Node[T]{Data: k}
		l.tail = l.head
	} else {
		l.tail.Next = &Node[T]{Data: k}
		l.tail = l.tail.Next
	}
	l.size++
}

// Head implements LinkedList.
func (l *linkedList[T]) Head() *Node[T] {
	if l.head == nil {
		return nil
	}
	return l.head
}

// IsEmpty implements LinkedList.
func (l *linkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Iterator implements LinkedList.
func (l *linkedList[T]) Iterator() Iterator[*Node[T]] {
	return newLinkedListIterator(l.head)
}

// Prepend implements LinkedList.
func (l *linkedList[T]) Prepend(k T) {
	if l.head == nil {
		l.head = &Node[T]{Data: k}
		l.tail = l.head
	} else {
		newHead := &Node[T]{Data: k, Next: l.head}
		l.head = newHead
	}
	l.size++
}

// Remove implements LinkedList.
func (l *linkedList[T]) Remove(node *Node[T]) {
	if node == nil {
		return
	}
	if l.head == node {
		l.head = node.Next
	} else {
		prev := l.head
		for prev != nil && prev.Next != node {
			prev = prev.Next
		}
		if prev != nil {
			prev.Next = node.Next
		}
	}
	if l.tail == node {
		l.tail = nil
	}
	l.size--
	if l.size == 0 {
		l.head = nil
		l.tail = nil
	}
	node.Next = nil
	node = nil
}

// Size implements LinkedList.
func (l *linkedList[T]) Size() int {
	return l.size
}

// Tail implements LinkedList.
func (l *linkedList[T]) Tail() *Node[T] {
	return l.tail
}

func NewLinkedList[T any]() LinkedList[T] {
	return &linkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}
