package main

import (
	"fmt"
	"strings"
)

// A Set of values of type T
type Set[T comparable] map[T]struct{}

// NewSet creates a new set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Has returns if the set has the given value
func (s Set[T]) Has(value T) bool {
	_, exists := s[value]
	return exists
}

// Add adds values to s
func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

// Remove removes values from s
func (s Set[T]) Remove(values ...T) {
	for _, v := range values {
		delete(s, v)
	}
}

func (s Set[T]) String() string {
	builder := strings.Builder{}
	for v := range s {
		if builder.Len() > 0 {
			builder.WriteString(", ")
		}
		fmt.Fprint(&builder, v)
	}
	return builder.String()
}

func main() {
	s := NewSet[string]()
	s.Add("a", "b", "c")
	fmt.Println(s)
	fmt.Println("s has a?", s.Has("a"))
	fmt.Println("s has d?", s.Has("d"))
}
