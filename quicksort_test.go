package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPartition(t *testing.T) {
	arr := []int{3, 4, 0, 2}
	QuicksortPartition(arr, 0, len(arr)-1)
	require.Equal(t, []int{0, 3, 4, 2}, arr)

}

func TestQuicksort(t *testing.T) {
	arr := []int{3, 2, 1, 7, 4, 9, 0, 1, 2, 3, -1, -4, -2}

	Quicksort(arr, 0, len(arr)-1)

	t.Log(arr)
	require.Equal(t, []int{-4, -2, -1, 0, 1, 1, 2, 2, 3, 3, 4, 7, 9}, arr)
}
