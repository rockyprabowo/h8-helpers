package slices

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"reflect"
	"testing"
)

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
			if gotDiff := Diff(tt.args.a, tt.args.b); !slices.Equal(gotDiff, tt.wantDiff) {
				t.Errorf("Diff() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func TestSliceLongShort(t *testing.T) {
	longSlice := []int{1, 3, 4, 0, 12, 9}
	shortSlice := []int{5, 1, 2, 9}
	shortSecondSlice := []int{4, 123, 56, 19}
	type args[T any] struct {
		a []T
		b []T
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantLong  []T
		wantShort []T
	}
	tests := []testCase[int]{
		{
			name: "should return longer slice on the first return value and shorter one on the second",
			args: args[int]{
				a: longSlice,
				b: shortSlice,
			},
			wantLong:  longSlice,
			wantShort: shortSlice,
		},
		{
			name: "should be able to return the same order if both array are the same length",
			args: args[int]{
				a: shortSlice,
				b: shortSecondSlice,
			},
			wantLong:  shortSlice,
			wantShort: shortSecondSlice,
		},
		{
			name: "should return longer slice on the first return value and shorter one on the second",
			args: args[int]{
				a: shortSlice,
				b: longSlice,
			},
			wantLong:  longSlice,
			wantShort: shortSlice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLong, gotShort := SliceLongShort(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(gotLong, tt.wantLong) {
				t.Errorf("SliceLongShort() gotLong = %v, want %v", gotLong, tt.wantLong)
			}
			if !reflect.DeepEqual(gotShort, tt.wantShort) {
				t.Errorf("SliceLongShort() gotShort = %v, want %v", gotShort, tt.wantShort)
			}
		})
	}
}
