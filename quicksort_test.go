package sorting

import (
	"math/rand"
	"testing"

	"golang.org/x/exp/constraints"
)

func init() {
	debugLog = true
}

func verifyPartition[T constraints.Ordered](t *testing.T, pivot T, arr []T) {
	lessSide := true
	for _, v := range arr {
		if lessSide {
			if v >= pivot {
				lessSide = false
			}
		} else {
			if v < pivot {
				t.Fatalf("right partition is not all > than pivot (%v < %v) %v", v, pivot, arr)
			}
		}
	}
}

func testPart[T constraints.Ordered](t *testing.T, inp []T) {
	pivot := inp[0]
	QuicksortPartition(inp, pivot, 0, len(inp)-1)
	verifyPartition(t, pivot, inp)
}

func TestPartition(t *testing.T) {
	testPart(t, []int{3, 4, 0, 2})
	testPart(t, []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2})
}

func FuzzPartition(f *testing.F) {
	f.Add(int64(20))
	f.Fuzz(func(t *testing.T, seed int64) {
		source := rand.NewSource(seed)

		inp := make([]int, source.Int63()%20)
		for i := 0; i < len(inp); i++ {
			inp[i] = int(source.Int63() % 20)
		}
		testPart(t, inp)
	})
}

// func TestQuicksort(t *testing.T) {
// 	arr := []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2}
// 	Quicksort(arr, 0, len(arr)-1)
// 	t.Log(arr)
// 	require.Equal(t, []int{-4, -2, -1, 0, 1, 1, 2, 2, 3, 3, 4, 7, 9}, arr)
// }
