package main

import (
	"testing"
)

func BenchmarkEqual4(b *testing.B) {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}

	Equal(x, y)
}

func BenchmarkEqualWithReflect4(b *testing.B) {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}

	EqualWithReflect(x, y)
}

func BenchmarkEqual10000000(b *testing.B) {
	x := make([]int, 10000000)
	y := make([]int, 10000000)

	Equal(x, y)
}

func BenchmarkEqualWithReflect10000000(b *testing.B) {
	x := make([]int, 10000000)
	y := make([]int, 10000000)

	EqualWithReflect(x, y)
}

func BenchmarkEqualDifferentElements(b *testing.B) {
	x := make([]int, 10000000)
	y := make([]int, 10000000)

	x[2] = 1

	Equal(x, y)
}

func BenchmarkEqualWithReflectDifferentElements(b *testing.B) {
	x := make([]int, 10000000)
	y := make([]int, 10000000)

	x[2] = 1

	EqualWithReflect(x, y)
}

func BenchmarkEqualDifferentLength(b *testing.B) {
	x := make([]int, 10000000)
	y := make([]int, 20000000)

	Equal(x, y)
}

func BenchmarkEqualWithReflectDifferentLength(b *testing.B) {
	x := make([]int, 10000000)
	y := make([]int, 20000000)

	EqualWithReflect(x, y)
}
