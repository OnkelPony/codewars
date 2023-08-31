package kata_test

import (
	"kata"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumDigPow(t *testing.T) {
	require.Equal(t, []uint64{13, 14, 15, 16, 17, 18, 19, 20}, SumDigPow(13, 20))
	require.Equal(t, []uint64{13}, SumDigPow(13, 13))
	require.Equal(t, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}, SumDigPow(1, 9))
}
