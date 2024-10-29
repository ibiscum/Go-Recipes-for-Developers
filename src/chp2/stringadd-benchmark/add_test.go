package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkXPlusY(b *testing.B) {
	x := "Hello"
	y := " World"
	for i := 0; i < b.N; i++ {
		_ = x + y
	}
}

func BenchmarkSprintf(b *testing.B) {
	x := "Hello"
	y := " World"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", x, y)
	}
}

func BenchmarkBuilder(b *testing.B) {
	x := "Hello"
	y := " World"
	for i := 0; i < b.N; i++ {
		builder := strings.Builder{}
		builder.WriteString(x)
		builder.WriteString(y)
	}
}
