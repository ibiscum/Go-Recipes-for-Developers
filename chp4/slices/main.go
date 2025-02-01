package main

import (
	"fmt"
	"slices"
)

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (val T) {
	val = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func main() {
	// Create an empty integer slice
	islice := make([]int, 0)
	// Append values 1, 2, 3 to islice, assign it to newSlice
	newSlice := append(islice, 1, 2, 3)
	fmt.Println(islice)
	fmt.Println(newSlice)

	// Create an empty integer slice
	islice = make([]int, 0)
	// Another integer slice with 3 elements
	otherSlice := []int{1, 2, 3}
	// Append 'otherSlice' to 'islice'
	newSlice = append(islice, otherSlice...)
	// Append 'otherSlice' to 'newSlice
	newSlice = append(newSlice, otherSlice...)
	fmt.Println(newSlice)
	fmt.Println(islice)

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice, cap(slice))
	suffix := slice[1:]
	fmt.Println(suffix)
	suffix2 := slice[3:]
	fmt.Println(suffix2)
	prefix := slice[:5]
	fmt.Println(prefix)

	mid := slice[3:6]
	fmt.Println(mid)

	edges := slices.Delete(slice, 3, 7)
	fmt.Println(edges)
	fmt.Println(slice)
	inserted := slices.Insert(slice, 3, 3, 4)
	fmt.Println(inserted)
	fmt.Println(edges)
	fmt.Println(slice)

	slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Keep two indexes, one to read from, one to write to
	write := 0
	for _, elem := range slice {
		if elem%2 == 0 { // Copy onle even numbers
			slice[write] = elem
			write++
		}
	}
	// Truncate the slice
	slice = slice[:write]
	fmt.Println(slice)

	intStack := Stack[int]{}
	intStack.Push(1)
	fmt.Println(intStack)
	fmt.Println(intStack.Pop())
	fmt.Println(intStack)
}
