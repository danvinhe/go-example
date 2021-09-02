package sortedslice

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSortedSlice(t *testing.T) {
	s := New(func(a, b interface{}) bool {
		if a.(int) >= b.(int) {
			return false
		}
		return true
	})
	tests := []struct {
		name  string
		value interface{}
		want  int
	}{
		{
			name:  "Add 20",
			value: 20,
			want:  0,
		},
		{
			name:  "Add 10",
			value: 10,
			want:  0,
		},
		{
			name:  "Add 30",
			value: 30,
			want:  2,
		},
		{
			name:  "Again Add 10",
			value: 10,
			want:  0,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Add(tt.value); got != tt.want {
				t.Errorf("SortedSlice.Add() = %v, want %v", got, tt.want)
			}
			if got := s.Len(); got != i+1 {
				t.Errorf("SortedSlice.Len() = %v, want %v", got, i+1)
			}
			t.Log(s.String())
		})
	}

	index := rand.Intn(len(tests))
	t.Run("Remove", func(t *testing.T) {
		if got := s.Remove(index); got != true {
			t.Errorf("SortedSlice.Add() = %v, want %v", got, true)
		}
		t.Log(s.String())
	})

	t.Run("Clear", func(t *testing.T) {
		s.Clear()
		t.Log(s.String())
	})
}
