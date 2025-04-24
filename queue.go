package ds

type Queue[T any] interface {
	Enqueue(k T)
	Dequeue() T
	Front() T
	IsEmpty() bool
	Size() int
}

type queue[T any] struct {
	linedList LinkedList[T]
}

// Dequeue implements Queue.
func (q *queue[T]) Dequeue() T {
	if q.linedList.IsEmpty() {
		panic("no objects in queue")
	}
	head := q.linedList.Head()
	q.linedList.Remove(head)
	return head.Data
}

// Enqueue implements Queue.
func (q *queue[T]) Enqueue(k T) {
	q.linedList.Append(k)
}

// Front implements Queue.
func (q *queue[T]) Front() T {
	if q.linedList.IsEmpty() {
		panic("no objects in queue")
	}
	head := q.linedList.Head()
	return head.Data
}

// IsEmpty implements Queue.
func (q *queue[T]) IsEmpty() bool {
	return q.linedList.IsEmpty()
}

// Size implements Queue.
func (q *queue[T]) Size() int {
	return q.linedList.Size()
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{
		linedList: NewLinkedList[T](),
	}
}
