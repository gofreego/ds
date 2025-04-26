package ds

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"NewStack"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack[int]()
			if s.IsEmpty() != true {
				t.Errorf("NewStack() = %v, want %v", s.IsEmpty(), true)
			}
			s.Push(1)
			if s.IsEmpty() != false {
				t.Errorf("NewStack() = %v, want %v", s.IsEmpty(), false)
			}
			if s.Top() != 1 {
				t.Errorf("NewStack() = %v, want %v", s.Top(), 1)
			}
			s.Push(2)
			if s.Top() != 2 {
				t.Errorf("NewStack() = %v, want %v", s.Top(), 2)
			}
			if s.Pop() != 2 {
				t.Errorf("NewStack() = %v, want %v", s.Pop(), 2)
			}
			if s.Top() != 1 {
				t.Errorf("NewStack() = %v, want %v", s.Top(), 1)
			}
			if s.Pop() != 1 {
				t.Errorf("NewStack() = %v, want %v", s.Pop(), 1)
			}
			if s.IsEmpty() != true {
				t.Errorf("NewStack() = %v, want %v", s.IsEmpty(), true)
			}
		})
	}
}
