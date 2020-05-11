package snapshot_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEntryIsDeleted(t *testing.T) {
	e := entry{
		verAdd: 1,
	}
	for i := 0; i < 2; i++ {
		require.False(t, e.isDeleted(version(i)))
	}
	e.verDel = 3
	require.False(t, e.isDeleted(2))
	require.True(t, e.isDeleted(3))
	require.True(t, e.isDeleted(4))
}

func TestAddAndDelete(t *testing.T) {
	arr := NewSnapshotArrayList()
	for i := 0; i < 5; i++ {
		arr.Add(i)
		require.Equal(t, arr.data[i].verAdd, version(i+1))
	}
	arr.Remove(4)
	arr.Remove(2)
	require.Equal(t, arr.data[4].verDel, version(6))
	require.Equal(t, arr.data[2].verDel, version(7))
	expectedLatest := make([]int, 0)
	for _, item := range arr.latestData {
		i, _ := item.val.(int)
		expectedLatest = append(expectedLatest, i)
	}
	require.Equal(t, expectedLatest, []int{0, 1, 3})
}
