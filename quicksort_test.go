package sorting

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func init() {
	debugLog = true
}

func verifyPartition[T constraints.Ordered](t *testing.T, pivot T, arr []T) {
	pivotIdx := 0
	for i, v := range arr {
		if v == pivot {
			pivotIdx = i
			break
		}
	}

	for i, v := range arr {
		if i < pivotIdx {
			require.Less(t, v, pivot, "left partition not less than pivot")
		}
		if i > pivotIdx {
			require.Greater(t, v, pivot, "right partition not greater than pivot")
		}
	}
}

func TestPartition(t *testing.T) {
	arr := []int{3, 4, 0, 2}
	pivotLoc := QuicksortPartition(arr, 0, len(arr)-1)
	verifyPartition(t, pivotLoc, arr)

	arr = []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2}
	pivotLoc = QuicksortPartition(arr, 0, len(arr)-1)
	verifyPartition(t, pivotLoc, arr)
}

func FuzzPartition(f *testing.F) {
	f.Add(int64(20))
	f.Fuzz(func(t *testing.T, seed int64) {
		source := rand.NewSource(seed)

		inp := make([]int, 7)
		for i := 0; i < len(inp); i++ {
			inp[i] = int(source.Int63())
		}

		pivot := inp[getPivotIdx(0, len(inp)-1)]
		QuicksortPartition(inp, 0, len(inp)-1)
		verifyPartition(t, pivot, inp)

		inp = make([]int, 8)
		for i := 0; i < len(inp); i++ {
			inp[i] = int(source.Int63())
		}

		pivot = inp[getPivotIdx(0, len(inp)-1)]
		QuicksortPartition(inp, 0, len(inp)-1)
		verifyPartition(t, pivot, inp)
	})
}

// func TestQuicksort(t *testing.T) {
// 	arr := []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2}
// 	Quicksort(arr, 0, len(arr)-1)
// 	t.Log(arr)
// 	require.Equal(t, []int{-4, -2, -1, 0, 1, 1, 2, 2, 3, 3, 4, 7, 9}, arr)
// }
