package slices

import (
	"golang.org/x/exp/constraints"
	"slices"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a []T
		b []T
	}
	type testCase[T constraints.Ordered] struct {
		name     string
		args     args[T]
		wantDiff []T
	}
	tests := []testCase[int]{
		{
			name: "should return the intersections between a and b with any duplicates",
			args: args[int]{
				a: []int{1, 3, 5, 21, 23, 99, 5, 4},
				b: []int{3, 22, 100, 5, 21, 1},
			},
			wantDiff: []int{1, 3, 5, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := Intersect(tt.args.a, tt.args.b); !slices.Equal(gotDiff, tt.wantDiff) {
				t.Errorf("Intersect() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}
