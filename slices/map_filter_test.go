package slices

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

type Person struct {
	name string
	age  uint
}

func TestMapExample(t *testing.T) {

	personList := []Person{
		{name: "Rocky", age: 25},
		{name: "Prabowo", age: 52},
	}
	personNames := Map(personList, func(p Person) string {
		return p.name
	})

	personAges := Map(personList, func(p Person) uint {
		return p.age
	})

	fmt.Println(personNames)
	fmt.Println(personAges)
}

func TestMap(t *testing.T) {
	type args[T any, U constraints.Unsigned] struct {
		slice []T
		f     func(T) U
	}

	tests := []struct {
		name       string
		args       args[Person, uint]
		wantMapped []uint
	}{
		{
			name: "should map all of the age from slice of Person",
			args: args[Person, uint]{
				slice: []Person{
					{name: "Rocky", age: 25},
					{name: "Prabowo", age: 21},
				},
				f: func(p Person) uint {
					return p.age
				},
			},
			wantMapped: []uint{
				25, 21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMapped := Map(tt.args.slice, tt.args.f); !reflect.DeepEqual(gotMapped, tt.wantMapped) {
				t.Errorf("Map() = %v, want %v", gotMapped, tt.wantMapped)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		slice []T
		f     func(T) bool
	}
	tests := []struct {
		name         string
		args         args[Person]
		wantFiltered []Person
	}{
		{
			name: "should return every Person with age > 20 from slice of Person",
			args: args[Person]{
				slice: []Person{
					{name: "Rocky", age: 15},
					{name: "Prabowo", age: 25},
					{name: "Alpha", age: 19},
					{name: "Beta", age: 20},
					{name: "Charlie", age: 23},
				},
				f: func(p Person) bool {
					return p.age > 20
				},
			},
			wantFiltered: []Person{
				{name: "Prabowo", age: 25},
				{name: "Charlie", age: 23},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFiltered := Filter(tt.args.slice, tt.args.f); !reflect.DeepEqual(gotFiltered, tt.wantFiltered) {
				t.Errorf("Filter() = %v, want %v", gotFiltered, tt.wantFiltered)
			}
		})
	}
}
