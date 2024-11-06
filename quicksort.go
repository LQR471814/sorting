package sorting

import (
	"fmt"
	"strings"

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

func getPivotIdx(firstIdx, lastIdx int) int {
	return (firstIdx + lastIdx) / 2
}

func QuicksortPartition[T constraints.Ordered](arr []T, firstIdx, lastIdx int) int {
	if len(arr) <= 2 {
		return -1
	}

	pivot := arr[getPivotIdx(firstIdx, lastIdx)]
	left := firstIdx - 1
	right := lastIdx + 1

	if debugLog {
		fmt.Printf("PARTITION (pivot: %v, left_idx: %d, right_idx: %d)\n", pivot, firstIdx, lastIdx)
	}

	testIndices(arr, firstIdx, lastIdx)

	for {
		for {
			left++
			if arr[left] >= pivot {
				break
			}
		}
		for {
			right--
			if arr[right] <= pivot {
				break
			}
		}

		if debugLog {
			lens := make([]int, lastIdx-firstIdx+1)
			for i := firstIdx; i <= lastIdx; i++ {
				out := fmt.Sprint(arr[i]) + " "
				lens[i] = len(out)
				fmt.Print(out)
			}
			fmt.Println()
			for i := firstIdx; i <= lastIdx; i++ {
				if i == left {
					fmt.Print("L" + strings.Repeat(" ", lens[i]-1))
					continue
				}
				if i == right {
					fmt.Print("R" + strings.Repeat(" ", lens[i]-1))
					continue
				}
				fmt.Print(strings.Repeat(" ", lens[i]))
			}
			fmt.Println()
		}

		if left >= right {
			if debugLog {
				fmt.Println("PARTITION RESULT", arr, right)
			}

			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
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
