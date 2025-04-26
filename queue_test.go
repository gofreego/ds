package ds

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"NewQueue"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()
			if q.IsEmpty() != true {
				t.Errorf("NewQueue() = %v, want %v", q.IsEmpty(), true)
			}
			if q.Size() != 0 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 0)
			}
			q.Enqueue(1)
			if q.IsEmpty() != false {
				t.Errorf("NewQueue() = %v, want %v", q.IsEmpty(), false)
			}
			if q.Size() != 1 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 1)
			}
			if q.Front() != 1 {
				t.Errorf("NewQueue() = %v, want %v", q.Front(), 1)
			}
			q.Enqueue(2)
			if q.Size() != 2 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 2)
			}
			if q.Front() != 1 {
				t.Errorf("NewQueue() = %v, want %v", q.Front(), 1)
			}
			if q.Dequeue() != 1 {
				t.Errorf("NewQueue() = %v, want %v", q.Dequeue(), 1)
			}
			if q.Size() != 1 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 1)
			}
			if q.Front() != 2 {
				t.Errorf("NewQueue() = %v, want %v", q.Front(), 2)
			}
			if q.Dequeue() != 2 {
				t.Errorf("NewQueue() = %v, want %v", q.Dequeue(), 2)
			}
			if q.IsEmpty() != true {
				t.Errorf("NewQueue() = %v, want %v", q.IsEmpty(), true)
			}
			if q.Size() != 0 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 0)
			}
			q.Enqueue(3)
			q.Enqueue(4)
			q.Enqueue(5)
			if q.Size() != 3 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 3)
			}
			if q.Front() != 3 {
				t.Errorf("NewQueue() = %v, want %v", q.Front(), 3)
			}
			if q.Dequeue() != 3 {
				t.Errorf("NewQueue() = %v, want %v", q.Dequeue(), 3)
			}
			if q.Size() != 2 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 2)
			}
			if q.Front() != 4 {
				t.Errorf("NewQueue() = %v, want %v", q.Front(), 4)
			}
			if q.Dequeue() != 4 {
				t.Errorf("NewQueue() = %v, want %v", q.Dequeue(), 4)
			}
			if q.Size() != 1 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 1)
			}
			if q.Front() != 5 {
				t.Errorf("NewQueue() = %v, want %v", q.Front(), 5)
			}
			if q.Dequeue() != 5 {
				t.Errorf("NewQueue() = %v, want %v", q.Dequeue(), 5)
			}
			if q.IsEmpty() != true {
				t.Errorf("NewQueue() = %v, want %v", q.IsEmpty(), true)
			}
			if q.Size() != 0 {
				t.Errorf("NewQueue() = %v, want %v", q.Size(), 0)
			}
		})
	}
}
