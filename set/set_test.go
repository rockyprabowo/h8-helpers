package set

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"slices"
	"testing"
)

func TestNewSet(t *testing.T) {
	expectedFixture := NewSet[string]()
	type testCase[T constraints.Ordered] struct {
		name string
		want *Set[T]
	}
	tests := []testCase[string]{
		{
			name: "should return a new, empty set",
			want: expectedFixture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet[string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSetFromSlice(t *testing.T) {
	sliceFixture := []string{
		"hello",
		"bonkers",
	}
	fixture := NewSetFromSlice[string](sliceFixture)

	sliceFixtureWithDups := []string{
		"hello",
		"bonkers",
		"hello",
		"hello",
		"bonkers",
		"world",
	}
	expectedSliceFixtureWithDups := []string{
		"hello",
		"bonkers",
		"world",
	}
	expectedFixtureWithDups := NewSetFromSlice(expectedSliceFixtureWithDups)
	type args[T constraints.Ordered] struct {
		slice []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want *Set[T]
	}
	tests := []testCase[string]{
		{
			name: "should create a new slice with contents",
			args: args[string]{
				slice: sliceFixture,
			},
			want: fixture,
		},
		{
			name: "should create a deduplicated set from slice with duplicate contents",
			args: args[string]{
				slice: sliceFixtureWithDups,
			},
			want: expectedFixtureWithDups,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetFromSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	inputFixture := "added"
	fixture := NewSet[string]()
	expectedFixture := NewSetFromSlice(
		[]string{
			inputFixture,
		})
	type args[T constraints.Ordered] struct {
		v T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{
			name: "should be able to add value to set",
			s:    *fixture,
			args: args[string]{
				v: inputFixture,
			},
			want: *expectedFixture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.v)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_AddMany(t *testing.T) {
	inputFixture := []string{
		"hello",
		"bonkers",
	}
	expectedFixture := NewSetFromSlice[string](inputFixture)

	inputFixtureWithDups := []string{
		"hello",
		"bonkers",
		"hello",
		"hello",
		"bonkers",
		"world",
	}
	expectedSliceFixtureWithDups := []string{
		"hello",
		"bonkers",
		"world",
	}
	expectedFixtureWithDups := NewSetFromSlice(expectedSliceFixtureWithDups)
	type args[T constraints.Ordered] struct {
		v []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{
			name: "should be able to add value to set",
			s:    *expectedFixture,
			args: args[string]{
				v: inputFixture,
			},
			want: *expectedFixture,
		},
		{
			name: "should be able to add duplicated values to set and deduplicated",
			s:    *expectedFixture,
			args: args[string]{
				v: inputFixtureWithDups,
			},
			want: *expectedFixtureWithDups,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddMany(tt.args.v...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("AddMany() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	inputFixture := []string{
		"hello",
		"bonkers",
	}
	targetFixture := NewSetFromSlice[string](inputFixture)
	expectedFixture := NewSet[string]()

	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{
			name: "should be able to clear the set",
			s:    *targetFixture,
			want: *expectedFixture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Clear() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Has(t *testing.T) {
	const happyPathValue = "hello"
	const badPathValue = "world"
	inputSliceFixture := []string{
		happyPathValue,
		"bonkers",
	}
	inputFixture := NewSetFromSlice[string](inputSliceFixture)

	type args[T constraints.Ordered] struct {
		v T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "should be able to return true if the value exists",
			s:    *inputFixture,
			args: args[string]{
				v: happyPathValue,
			},
			want: true,
		},
		{
			name: "should be able to return false if the value doesn't exists",
			s:    *inputFixture,
			args: args[string]{
				v: badPathValue,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Has(tt.args.v); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	const happyPathValue = "hello"
	inputSliceFixture := []string{
		happyPathValue,
		"bonkers",
		"world",
	}
	inputFixture := NewSetFromSlice[string](inputSliceFixture)

	expectedInputSliceFixture := []string{
		"bonkers",
		"world",
	}
	expectedFixture := NewSetFromSlice(expectedInputSliceFixture)

	type args[T constraints.Ordered] struct {
		v T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{
			name: "should be able to remove a value from set",
			s:    *inputFixture,
			args: args[string]{
				happyPathValue,
			},
			want: *expectedFixture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.v)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Remove() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_RemoveMany(t *testing.T) {
	happyPathValues := []string{"hello", "world"}
	inputSliceFixture := append(happyPathValues, "bonkers")
	inputFixture := NewSetFromSlice[string](inputSliceFixture)

	expectedInputSliceFixture := []string{
		"bonkers",
	}
	expectedFixture := NewSetFromSlice(expectedInputSliceFixture)
	type args[T constraints.Ordered] struct {
		v []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{
			name: "Should be able to remove many values from the set",
			s:    *inputFixture,
			args: args[string]{
				v: happyPathValues,
			},
			want: *expectedFixture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.RemoveMany(tt.args.v...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("RemoveMany() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	inputFixture := []string{
		"hello",
		"bonkers",
	}
	targetFixture := NewSetFromSlice[string](inputFixture)
	expectedValue := len(inputFixture)

	inputFixtureWithDups := []string{
		"hello",
		"bonkers",
		"hello",
		"hello",
		"bonkers",
		"world",
	}
	expectedSliceFixtureWithDups := []string{
		"hello",
		"bonkers",
		"world",
	}
	targetWithDupsFixture := NewSetFromSlice(inputFixtureWithDups)
	expectedWithDupsValue := len(expectedSliceFixtureWithDups)

	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		want int
	}
	tests := []testCase[string]{
		{
			name: "should be able to return a valid size of set",
			s:    *targetFixture,
			want: expectedValue,
		},
		{
			name: "should be able to return a valid size of deduplicated set",
			s:    *targetWithDupsFixture,
			want: expectedWithDupsValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToSlice(t *testing.T) {
	inputFixture := []string{
		"hello",
		"bonkers",
	}
	targetFixture := NewSetFromSlice[string](inputFixture)
	expectedResultFixture := []string{
		"bonkers",
		"hello",
	}
	type testCase[T constraints.Ordered] struct {
		name string
		s    Set[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "should be able to return a sorted slice form of set",
			s:    *targetFixture,
			want: expectedResultFixture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToSlice(); !slices.Equal(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
