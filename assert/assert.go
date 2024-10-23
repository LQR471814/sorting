package assert

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Equal[T comparable](a, b T, failureMessage string) {
	if a != b {
		panic(fmt.Sprintf("%s: %v != %v", failureMessage, a, b))
	}
}

func Gt[T constraints.Ordered](a, b T, failureMessage string) {
	if a <= b {
		panic(fmt.Sprintf("%s: %v !> %v", failureMessage, a, b))
	}
}

func Gte[T constraints.Ordered](a, b T, failureMessage string) {
	if a < b {
		panic(fmt.Sprintf("%s: %v !>= %v", failureMessage, a, b))
	}
}

func Lt[T constraints.Ordered](a, b T, failureMessage string) {
	if a >= b {
		panic(fmt.Sprintf("%s: %v !< %v", failureMessage, a, b))
	}
}

func Lte[T constraints.Ordered](a, b T, failureMessage string) {
	if a > b {
		panic(fmt.Sprintf("%s: %v !<= %v", failureMessage, a, b))
	}
}
