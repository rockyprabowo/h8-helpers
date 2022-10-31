package slices

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestDiffWithDuplicates(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a []T
		b []T
	}
	tests := []struct {
		name     string
		args     args[uint]
		wantDiff []uint
	}{
		{
			name: "should return the difference between a and b with any duplicates",
			args: args[uint]{
				a: []uint{1, 2, 3, 3, 3, 4, 5, 6},
				b: []uint{1, 2, 3},
			},
			wantDiff: []uint{3, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := DiffWithDuplicates(tt.args.a, tt.args.b); !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("DiffWithDuplicates() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a []T
		b []T
	}
	tests := []struct {
		name     string
		args     args[uint]
		wantDiff []uint
	}{
		{
			name: "should return the difference between a and b",
			args: args[uint]{
				a: []uint{1, 2, 3, 3, 3, 4, 5, 6},
				b: []uint{1, 2, 3},
			},
			wantDiff: []uint{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := Diff(tt.args.a, tt.args.b); !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("Diff() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}
