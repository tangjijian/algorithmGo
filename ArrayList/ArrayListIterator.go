package ArrayList

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove()
	GetIndex() int
}

type ArrayIterator struct {
	currentIndex int
	list         *ArrayList
}

func (it *ArrayIterator) HasNext() bool {
	return it.list.TheSize > it.currentIndex
}

func (it *ArrayIterator) Next() interface{} {
	if it.HasNext() {
		val := it.list.dataStore[it.currentIndex]
		it.currentIndex++
		return val
	}
	return nil
}
func (it *ArrayIterator) Remove() {
	_ = it.list.Delete(it.currentIndex)
	it.currentIndex--
}
func (it *ArrayIterator) GetIndex() int {
	return 1
}
