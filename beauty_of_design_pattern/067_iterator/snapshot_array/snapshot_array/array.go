package snapshot_array

type version uint64

type entry struct {
	val            interface{}
	verAdd, verDel version
}

func (e *entry) isDeleted(v version) bool {
	return e.verDel > 0 && v >= e.verDel
}

type SnapshotArrayList struct {
	latestData, data []*entry
	curVersion       version
}

func NewSnapshotArrayList() *SnapshotArrayList {
	return &SnapshotArrayList{
		latestData: make([]*entry, 0, 10),
		data:       make([]*entry, 0, 10),
	}
}

func (l *SnapshotArrayList) Add(obj interface{}) {
	l.curVersion++
	item := &entry{
		val:    obj,
		verAdd: l.curVersion,
	}
	l.data = append(l.data, item)
	l.latestData = append(l.latestData, item)
}

func (l *SnapshotArrayList) Remove(obj interface{}) {
	l.curVersion++
	// l.data 软删除
	for _, item := range l.data {
		if item.val == obj {
			item.verDel = l.curVersion
		}
	}
	// 重建l.latestData
	l.latestData = l.latestData[0:0]
	for _, item := range l.data {
		if !item.isDeleted(l.curVersion) {
			l.latestData = append(l.latestData, item)
		}
	}
}

func (l *SnapshotArrayList) Get(i int) interface{} {
	return l.latestData[i].val
}

func (l *SnapshotArrayList) Iterator() *SnapshotArrayIterator {
	return newSnapshotArrayIterator(l)
}
