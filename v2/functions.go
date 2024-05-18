package gomodule

import "golang.org/x/exp/constraints"

type Numeric interface {
	constraints.Integer | constraints.Float
}

// Add adds two int and returns the sum.
// Implements [Addition]
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Numeric](a T, b T) T {
	return a + b
}
