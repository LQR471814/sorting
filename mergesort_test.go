package sorting

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func testMergesort[T constraints.Ordered](t *testing.T, arr []T) {
	expected := make([]T, len(arr))
	copy(expected, arr)
	slices.Sort(expected)

	Mergesort(arr)
	require.Equal(t, expected, arr)
}

func FuzzMergesort(f *testing.F) {
	f.Add(int64(20))
	f.Fuzz(func(t *testing.T, seed int64) {
		source := rand.NewSource(seed)

		inp := make([]int, source.Int63()%20)
		for i := 0; i < len(inp); i++ {
			inp[i] = int(source.Int63() % 20)
		}
		testMergesort(t, inp)
	})
}

func TestMergesort(t *testing.T) {
	testMergesort(t, []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2})
}
