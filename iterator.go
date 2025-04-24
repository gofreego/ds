package ds

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}
