package ds

func Reduce[T any, U any](l []T, f func(T) U) []U {
	result := make([]U, len(l))
	for i, v := range l {
		result[i] = f(v)
	}
	return result
}

func ReduceAsync[T any, U any](l []T, f func(T) U) []U {
	result := make([]U, len(l))
	ch := make(chan struct {
		index int
		value U
	}, len(l))

	for i, v := range l {
		go func(i int, v T) {
			ch <- struct {
				index int
				value U
			}{i, f(v)}
		}(i, v)
	}

	for range l {
		r := <-ch
		result[r.index] = r.value
	}
	return result
}
