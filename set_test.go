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
				t.Errorf("NewSet() = %v, want nil", set.Values())
			}
			if set.Contains(1) {
				t.Errorf("NewSet() = %v, want false", set.Contains(1))
			}
			if set.Add(1) != true {
				t.Errorf("NewSet() = %v, want true", set.Add(1))
			}
			if set.Add(1) != false {
				t.Errorf("NewSet() = %v, want false", set.Add(1))
			}
			if set.Contains(1) != true {
				t.Errorf("NewSet() = %v, want true", set.Contains(1))
			}
			if set.Remove(1) != true {
				t.Errorf("NewSet() = %v, want true", set.Remove(1))
			}
			if set.Remove(1) != false {
				t.Errorf("NewSet() = %v, want false", set.Remove(1))
			}
			if set.Contains(1) != false {
				t.Errorf("NewSet() = %v, want false", set.Contains(1))
			}
			if set.Values() != nil {
				t.Errorf("NewSet() = %v, want nil", set.Values())
			}
			if set.Add(1) != true {
				t.Errorf("NewSet() = %v, want true", set.Add(1))
			}
			if set.Add(2) != true {
				t.Errorf("NewSet() = %v, want true", set.Add(2))
			}
			if set.Add(3) != true {
				t.Errorf("NewSet() = %v, want true", set.Add(3))
			}
			if set.Add(4) != true {
				t.Errorf("NewSet() = %v, want true", set.Add(4))
			}
			if set.Add(1) != false {
				t.Errorf("NewSet() = %v, want false", set.Add(1))
			}
			if set.Add(2) != false {
				t.Errorf("NewSet() = %v, want false", set.Add(2))
			}
			if set.Size() != 4 {
				t.Errorf("NewSet() = %v, want 4", set.Size())
			}
		})
	}
}
