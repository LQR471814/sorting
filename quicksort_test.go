package sorting

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func init() {
	debugLog = true
}

func verifyPartition[T constraints.Ordered](t *testing.T, pivotLoc int, arr []T) {
	for i, v := range arr {
		if i < pivotLoc {
			require.Less(t, v, arr[pivotLoc], "left partition not less than pivot")
		}
		if i > pivotLoc {
			require.Greater(t, v, arr[pivotLoc], "right partition not greater than pivot")
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
	source := rand.NewSource(123)

	f.Add(20)
	f.Fuzz(func(t *testing.T, n int) {
		n %= 20

		inp := make([]int, n)
		for i := 0; i < n; i++ {
			inp[i] = int(source.Int63())
		}

		fmt.Print("---- Run ----\n\n")

		pivot := QuicksortPartition(inp, 0, len(inp)-1)
		verifyPartition(t, pivot, inp)

		fmt.Print("\n---- End ----\n\n")
	})
}

// func TestQuicksort(t *testing.T) {
// 	arr := []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2}
// 	Quicksort(arr, 0, len(arr)-1)
// 	t.Log(arr)
// 	require.Equal(t, []int{-4, -2, -1, 0, 1, 1, 2, 2, 3, 3, 4, 7, 9}, arr)
// }
