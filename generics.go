package main

import "golang.org/x/exp/constraints"

func add[T constraints.Ordered](a, b T) T {
	return a + b
}
