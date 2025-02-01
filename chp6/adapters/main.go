package main

import (
	"fmt"
	"regexp"

	"golang.org/x/exp/constraints"
)

// ToPtr returns a pointer *T pointing to the given value
func ToPtr[T any](value T) *T {
	return &value
}

// ToSlice returns a slice []T{value}
func ToSlice[T any](value T) []T {
	return []T{value}
}

// SliceOf returns []T containing values
func SliceOf[T any](values ...T) []T {
	return values
}

type Number interface {
	constraints.Float | constraints.Integer
}

// ConvertNumberSlice converts a number slice to another type of a
// number slice
func ConvertNumberSlice[S, T Number](source []S) []T {
	result := make([]T, len(source))
	for i, v := range source {
		result[i] = T(v)
	}
	return result
}

// Last returns the last element of a slice, and whether or not that
// element exists
func Last[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[len(slice)-1], true
}

// Must returns value if there is no error, panics otherwise
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func main() {
	var x *int = ToPtr(1)
	fmt.Println("x = ToPtr(1)", x, *x)
	fmt.Println("ToSlice(1)", ToSlice(1))
	fmt.Println("ToSlice(str)", ToSlice("str"))
	fmt.Println("SliceOf(1,2,3)", SliceOf(1, 2, 3))
	fmt.Println("ConvertNumberSlice[float64, int]([]float64{1.1, 0.2, 33.41}))", ConvertNumberSlice[float64, int]([]float64{1.1, 0.2, 33.41}))
	v, ok := Last([]int{1, 2, 3, 4, 5})
	fmt.Println("Last([]int{1, 2, 3, 4, 5})", v, ok)
	v, ok = Last([]int{})
	fmt.Println("Last([]int{})", v, ok)
	fmt.Println("Must(regexp.Compile(test))", Must(regexp.Compile("test")))
}
