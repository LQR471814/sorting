package sorting

import (
	"fmt"

	"my-sorting/assert"

	"golang.org/x/exp/constraints"
)

var debugLog bool

func testIndices[T any](arr []T, indices ...int) {
	for i, idx := range indices {
		assert.Lt(idx, len(arr), fmt.Sprintf("idx #%d not < len(arr)", i))
		assert.Gte(idx, 0, fmt.Sprintf("idx #%d not >= 0", i))
	}
}

func QuicksortPartition[T constraints.Ordered](arr []T, firstIdx, lastIdx int) int {
	if len(arr) == 0 {
		return -1
	}
	if firstIdx == lastIdx {
		return -1
	}

	pivot := arr[(firstIdx+lastIdx)/2]
	left := firstIdx
	right := lastIdx

	testIndices(arr, firstIdx, lastIdx)

	for {
		if debugLog {
			for i := firstIdx; i <= lastIdx; i++ {
				fmt.Printf("%-3v", arr[i])
			}
			fmt.Println()
			for i := firstIdx; i <= lastIdx; i++ {
				if i == left {
					fmt.Print("L  ")
					continue
				}
				if i == right {
					fmt.Print("R  ")
					continue
				}
				fmt.Print("   ")
			}
			fmt.Println()
		}

		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left >= right {
			break
		}

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

	if debugLog {
		fmt.Println("PARTITION RESULT", arr, pivotIdx)
	}
	assert.Equal(arr[pivotIdx], pivot, "final idx value != pivot value")

	return pivotIdx
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
