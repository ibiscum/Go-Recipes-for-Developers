package main

import "fmt"

// Direction is an integer type
type Direction int

// Direction constants
const (
	DirectionLeft Direction = iota
	DirectionRight
)

// String returns a string representation of a direction
func (dir Direction) String() string {
	switch dir {
	case DirectionLeft:
		return "left"
	case DirectionRight:
		return "right"
	}
	return ""
}

func main() {
	fmt.Println(DirectionLeft, int(DirectionLeft))
	fmt.Println(DirectionRight, int(DirectionRight))
}
