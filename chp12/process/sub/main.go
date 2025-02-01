package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var max uint64
	max = math.MaxUint64
	if len(os.Args) > 1 {
		x, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		max = uint64(x)
	}
	for i := uint64(0); i < max; i++ {
		fmt.Println(i)
		// Print an error every 10 items
		if i%10 == 0 {
			fmt.Fprintln(os.Stderr, "Error:", i)
		}
	}
}
