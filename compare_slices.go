package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	fmt.Println(Equal(a, b))

	fmt.Println(EqualWithReflect(a, b))
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualWithReflect(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}

func EqualWithSlices(a, b []int) bool {
	return slices.Equal(a, b)
}
