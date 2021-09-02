package fibonacci

import (
	"reflect"
	"testing"
)

func TestSequence(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "max 0",
			args: args{0},
			want: []int{0},
		},
		{
			name: "max 1",
			args: args{1},
			want: []int{0, 1, 1},
		},
		{
			name: "max 2",
			args: args{2},
			want: []int{0, 1, 1, 2},
		},
		{
			name: "max 100",
			args: args{100},
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sequence(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{0},
			want: 0,
		},
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "2",
			args: args{2},
			want: 1,
		},
		{
			name: "10",
			args: args{10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := At(tt.args.index); got != tt.want {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}
		})
	}
}
