package sorting

import (
	"golang.org/x/exp/constraints"
)

func QuicksortPartition[T constraints.Ordered](arr []T, firstIdx, lastIdx int) int {
	pivot := arr[(firstIdx+lastIdx)/2]
	left := firstIdx
	right := lastIdx

	for left < right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}

		originalLeft := arr[left]
		arr[left] = arr[right]
		arr[right] = originalLeft
		left++
		right--
	}

	return right
}

func Quicksort[T constraints.Ordered](arr []T, firstIdx, lastIdx int) {
	if firstIdx >= lastIdx {
		return
	}

	newEnd := QuicksortPartition(arr, firstIdx, lastIdx)
	if newEnd < firstIdx {
		return
	}

	Quicksort(arr, firstIdx, newEnd)
	Quicksort(arr, newEnd+1, lastIdx)
}
