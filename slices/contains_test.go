package slices

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func TestContains(t *testing.T) {
	happyPathInput := "bonkers"
	badPathInput := "BONK"
	inputSlice := []string{
		"hello",
		happyPathInput,
		"bonkers",
	}

	type args[T constraints.Ordered] struct {
		slice  []T
		search T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "should return true when searched value in the input slice",
			args: args[string]{
				slice:  inputSlice,
				search: happyPathInput,
			},
			want: true,
		},
		{
			name: "should return false when searched value is not in the input slice",
			args: args[string]{
				slice:  inputSlice,
				search: badPathInput,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.search); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
