package ds

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

type listIterator[T any] struct {
	current []T
	index   int
}

func NewListIterator[T any](elements []T) Iterator[T] {
	return &listIterator[T]{current: elements, index: 0}
}

func (it *listIterator[T]) Next() T {
	if it.index >= len(it.current) {
		panic("no more elements")
	}
	element := it.current[it.index]
	it.index++
	return element
}

func (it *listIterator[T]) HasNext() bool {
	return it.index < len(it.current)
}
