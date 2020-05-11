package snapshot_array

type SnapshotArrayIterator struct {
	snapshotVersion        version
	cursorInAll, leftCount int
	arrayList              *SnapshotArrayList
}

func newSnapshotArrayIterator(list *SnapshotArrayList) *SnapshotArrayIterator {
	ret := &SnapshotArrayIterator{
		snapshotVersion: list.curVersion,
		cursorInAll:     0,
		leftCount:       len(list.latestData),
		arrayList:       list,
	}
	return ret
}

func (iter *SnapshotArrayIterator) HasNext() bool {
	return iter.leftCount > 0
}

func (iter *SnapshotArrayIterator) Next() interface{} {
	arr := iter.arrayList
	for {
		if !iter.HasNext() {
			return nil
		}
		if e := arr.data[iter.cursorInAll]; !e.isDeleted(iter.snapshotVersion) {
			iter.cursorInAll++
			iter.leftCount--
			return e.val
		}
		iter.cursorInAll++
	}
}
