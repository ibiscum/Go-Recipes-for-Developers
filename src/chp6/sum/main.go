package main

import "fmt"

// Sum1 accepts int and float64 values and returns the sum of those
func Sum1[T int | float64](values ...T) T {
	var result T
	for _, x := range values {
		result += x
	}
	return result
}

type number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
		~float64 | ~float32
}

func Sum2[T number](values ...T) T {
	var result T
	for _, x := range values {
		result += x
	}
	return result
}

func main() {
	// Type of Sum1 can be inferred from the arguments (int)
	fmt.Println(Sum1(1, 2, 3, 4, 5))
	// Type of Sum1 can be inferred from the arguments (float64)
	fmt.Println(Sum1(1.1, 12.5, 10.42))

	// Explicit specification of type. This version of Sum1 works with
	// float64 values.
	fmt.Println(Sum1[float64](1, 2, 3, 4, 5))

	// Sum2 works with all number types. Here, the result is uint16
	fmt.Println(Sum2(uint16(1), uint16(2), uint16(3)))

	// fun is a function that accepts float32 values
	fun := Sum2[float32]
	fmt.Println(fun(1, 2, 3))
}
