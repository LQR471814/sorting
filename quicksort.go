package sorting

import (
	"fmt"

	"my-sorting/assert"

	"golang.org/x/exp/constraints"
)

func testIndices[T any](arr []T, indices ...int) {
	for i, idx := range indices {
		assert.Lt(idx, len(arr), fmt.Sprintf("idx #%d not < len(arr)", i))
		assert.Gte(idx, 0, fmt.Sprintf("idx #%d not >= 0", i))
	}
}

func QuicksortPartition[T constraints.Ordered](arr []T, firstIdx, lastIdx int) int {
	pivot := arr[(firstIdx+lastIdx)/2]
	left := firstIdx
	right := lastIdx

	testIndices(arr, firstIdx, lastIdx)

	for {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left >= right {
			break
		}

		fmt.Printf("%v L[%d] R[%d] %v < %v; %v > %v\n", arr, left, right, arr[left], pivot, arr[right], pivot)

		originalLeft := arr[left]
		arr[left] = arr[right]
		arr[right] = originalLeft
		left++
		right--
	}

	pivotIdx := right
	if arr[pivotIdx] != pivot {
		pivotIdx = left
	}

	fmt.Println(arr, pivotIdx)
	assert.Equal(arr[pivotIdx], pivot, "final idx value != pivot value")

	return right + 1
}

func Quicksort[T constraints.Ordered](arr []T, firstIdx, lastIdx int) {
	if firstIdx >= lastIdx {
		return
	}

	testIndices(arr, firstIdx, lastIdx)

	newEnd := QuicksortPartition(arr, firstIdx, lastIdx)
	Quicksort(arr, firstIdx, newEnd-1)
	Quicksort(arr, newEnd+1, lastIdx)
}
