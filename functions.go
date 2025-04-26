package ds

import "sync"

func Map[T any, U any](l []T, f func(T) U) []U {
	result := make([]U, len(l))
	for i, v := range l {
		result[i] = f(v)
	}
	return result
}

func MapAsync[T any, U any](l []T, f func(T) U) []U {
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

func Filter[T any](l []T, f func(T) bool) []T {
	result := make([]T, 0, len(l))
	for _, v := range l {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterAsync[T any](l []T, f func(T) bool, concurrency ...int) []T {
	noOfGoroutines := 1
	if len(concurrency) > 0 {
		noOfGoroutines = concurrency[0]
	}
	ch := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(noOfGoroutines)
	result := make([]T, 0)

	for i := 0; i < noOfGoroutines; i++ {
		go func() {
			for v := range ch {
				if f(v) {
					result = append(result, v)
				}
			}
			wg.Done()
		}()
	}

	for _, v := range l {
		ch <- v
	}
	close(ch)
	wg.Wait()
	return result
}

func Reduce[T any, U any](l []T, f func(U, T) U, initial U) U {
	result := initial
	for _, v := range l {
		result = f(result, v)
	}
	return result
}

func ForEach[T any](l []T, f func(T)) {
	for _, v := range l {
		f(v)
	}
}

func ForEachAsync[T any](l []T, f func(T), concurrency ...int) {
	noOfGoroutines := 1
	if len(concurrency) > 0 {
		noOfGoroutines = concurrency[0]
	}
	ch := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(noOfGoroutines)
	for i := 0; i < noOfGoroutines; i++ {
		go func() {
			for v := range ch {
				f(v)
			}
			wg.Done()
		}()
	}

	for _, v := range l {
		ch <- v
	}
	close(ch)
	wg.Wait()
}
