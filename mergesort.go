package sorting

import (
	"golang.org/x/exp/constraints"
)

func Mergesort[T constraints.Ordered](arr []T) {
	mergesortInner(arr, make([]T, len(arr)))
}

func mergesortInner[T constraints.Ordered](arr []T, scratch []T) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	mergesortInner(arr[:mid], scratch)
	mergesortInner(arr[mid:], scratch)

	left := 0
	right := mid
	i := 0
	for left < mid || right < len(arr) {
		if left >= mid {
			scratch[i] = arr[right]
			right++
			i++
			continue
		}
		if right >= len(arr) {
			scratch[i] = arr[left]
			left++
			i++
			continue
		}

		leftVal, rightVal := arr[left], arr[right]
		if leftVal < rightVal {
			scratch[i] = arr[left]
			left++
			i++
		} else {
			scratch[i] = arr[right]
			right++
			i++
		}
	}

	for x := 0; x < len(arr); x++ {
		arr[x] = scratch[x]
	}
}
