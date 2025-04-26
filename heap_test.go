package ds

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	type args struct {
		t HeapType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "MinHeap",
			args: args{
				t: MinHeap,
			}},
		{
			name: "MaxHeap",
			args: args{
				t: MaxHeap,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.t == MinHeap {
				h := NewHeap(tt.args.t, func(a, b int) bool { return a < b })
				if !h.IsEmpty() {
					t.Errorf("Heap should be empty")
				}
				h.Push(1)
				if h.IsEmpty() {
					t.Errorf("Heap should not be empty")
				}
				if h.Size() != 1 {
					t.Errorf("Heap size should be 1")
				}
				if h.Top() != 1 {
					t.Errorf("Heap top should be 1")
				}
				h.Push(2)
				if h.Size() != 2 {
					t.Errorf("Heap size should be 2")
				}
				if h.Top() != 1 {
					t.Errorf("Heap top should be 1")
				}
				h.Push(3)
				if h.Size() != 3 {
					t.Errorf("Heap size should be 3")
				}
				if h.Top() != 1 {
					t.Errorf("Heap top should be 1")
				}
				h.Push(0)
				if h.Size() != 4 {
					t.Errorf("Heap size should be 4")
				}
				if h.Top() != 0 {
					t.Errorf("Heap top should be 0")
				}
				ele := h.Pop()
				if ele != 0 {
					t.Errorf("Heap pop should be 0")
				}
				if h.Size() != 3 {
					t.Errorf("Heap size should be 3")
				}
				if h.Top() != 1 {
					t.Errorf("Heap top should be 1")
				}
				ele = h.Pop()
				if ele != 1 {
					t.Errorf("Heap pop should be 1")
				}
				if h.Size() != 2 {
					t.Errorf("Heap size should be 2")
				}
				if h.Top() != 2 {
					t.Errorf("Heap top should be 2")
				}
				h.Pop()
				h.Pop()
				if h.Size() != 0 {
					t.Errorf("Heap size should be 0")
				}
				if !h.IsEmpty() {
					t.Errorf("Heap should be empty")
				}
			} else {
				h := NewHeap(tt.args.t, func(a, b int) bool { return a < b })
				if !h.IsEmpty() {
					t.Errorf("Heap should be empty")
				}
				h.Push(1)
				if h.IsEmpty() {
					t.Errorf("Heap should not be empty")
				}
				if h.Size() != 1 {
					t.Errorf("Heap size should be 1")
				}
				if h.Top() != 1 {
					t.Errorf("Heap top should be 1")
				}
				h.Push(2)
				if h.Size() != 2 {
					t.Errorf("Heap size should be 2")
				}
				if h.Top() != 2 {
					t.Errorf("Heap top should be 2")
				}
				h.Push(3)
				if h.Size() != 3 {
					t.Errorf("Heap size should be 3")
				}
				if h.Top() != 3 {
					t.Errorf("Heap top should be 3")
				}
				h.Push(0)
				if h.Size() != 4 {
					t.Errorf("Heap size should be 4")
				}
				if h.Top() != 3 {
					t.Errorf("Heap top should be 3")
				}
				ele := h.Pop()
				if ele != 3 {
					t.Errorf("Heap pop should be 3")
				}
				if h.Size() != 3 {
					t.Errorf("Heap size should be 3")
				}
				if h.Top() != 2 {
					t.Errorf("Heap top should be 2")
				}
				ele = h.Pop()
				if ele != 2 {
					t.Errorf("Heap pop should be 2")
				}
				if h.Size() != 2 {
					t.Errorf("Heap size should be 2")
				}
				if h.Push(1); h.Top() != 1 {
					t.Errorf("Heap top should be 1")
				}
				h.Push(1)
				h.Push(1)
				h.Push(5)
				h.Push(1)
				h.Push(5)
				h.Push(2)
				if h.Pop() != 5 {
					t.Errorf("Heap pop should be 5")
				}
				if h.Pop() != 5 {
					t.Errorf("Heap pop should be 5")
				}
				if h.Pop() != 2 {
					t.Errorf("Heap pop should be 2")
				}
				if h.Pop() != 1 {
					t.Errorf("Heap pop should be 1")
				}
				if h.Pop() != 1 {
					t.Errorf("Heap pop should be 1")
				}
				if h.Pop() != 1 {
					t.Errorf("Heap pop should be 1")
				}
				if h.Pop() != 1 {
					t.Errorf("Heap pop should be 1")
				}
				if ele := h.Pop(); ele != 1 {
					t.Errorf("Heap pop should be 0, got %d", ele)
				}

				if h.Pop() != 0 {
					t.Errorf("Heap pop should be 0")
				}

				if h.Size() != 0 {
					t.Errorf("Heap size should be 0, got %d", h.Size())
				}

			}
		})
	}
}
