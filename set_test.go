package ds

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet[int]()
			if set == nil {
				t.Errorf("NewSet() = nil, want not nil")
			}
			if reflect.TypeOf(set).String() != "*ds.set[int]" {
				t.Errorf("NewSet() = %T, want *ds.set[int]", set)
			}
			if set.Values() != nil {
				t.Errorf("set.Values() = %v, want nil", set.Values())
			}
			if set.Contains(1) {
				t.Errorf("set.Contains(1) = true, want false")
			}
			if set.Add(1) != true {
				t.Errorf("set.Add(1) = false, want true")
			}
			if set.Add(1) != false {
				t.Errorf("set.Add(1) = true, want false")
			}
			if set.Contains(1) != true {
				t.Errorf("set.Contains(1) = false, want true")
			}
			if set.Remove(1) != true {
				t.Errorf("set.Remove(1) = false, want true")
			}
			if set.Remove(1) != false {
				t.Errorf("set.Remove(1) = true, want false")
			}
			if set.Contains(1) != false {
				t.Errorf("set.Contains(1) = true, want false")
			}
			if set.Values() != nil {
				t.Errorf("set.Values() = %v, want nil", set.Values())
			}
			if set.Add(1) != true {
				t.Errorf("set.Add(1) = false, want true")
			}
			if set.Add(2) != true {
				t.Errorf("set.Add(2) = false, want true")
			}
			if set.Add(3) != true {
				t.Errorf("set.Add(3) = false, want true")
			}
			if set.Add(4) != true {
				t.Errorf("set.Add(4) = false, want true")
			}
			if set.Add(1) != false {
				t.Errorf("set.Add(1) = true, want false")
			}
			if set.Add(2) != false {
				t.Errorf("set.Add(2) = true, want false")
			}
			if set.Size() != 4 {
				t.Errorf("set.Size() = %d, want 4", set.Size())
			}
		})
	}
}
