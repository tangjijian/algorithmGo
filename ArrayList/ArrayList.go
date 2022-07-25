package ArrayList

import (
	"errors"
	"fmt"
)

type list interface {
	Get(index int) (interface{}, error)
	Set(index int, val interface{}) error
	Append(newVal interface{})
	Insert(index int, val interface{}) error
	Clear()
	Delete(index int) error
	Size() int
	String() string
	Iterator() Iterator
}

//type Iterable interface {
//	Iterator() Iterator
//}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayIterator)
	it.list = list
	it.currentIndex = 0
	return it
}

type ArrayList struct {
	TheSize   int
	dataStore []interface{}
}

func NewArrayList() list {
	list := new(ArrayList)
	list.TheSize = 0
	list.dataStore = make([]interface{}, 0, 10)
	return list
}
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("数组越界")
	}
	return list.dataStore[index], nil
}
func (list *ArrayList) Set(index int, val interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("数组越界")
	}
	list.dataStore[index] = val
	return nil
}
func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.TheSize++

}
func (list *ArrayList) Insert(index int, val interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("数组越界")
	}
	list.checkIsFull()
	list.dataStore = list.dataStore[:list.TheSize+1]
	for i := list.TheSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.TheSize++
	list.dataStore[index] = val
	return nil
}
func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStore) {
		newData := make([]interface{}, list.TheSize, 2*cap(list.dataStore))
		copy(newData, list.dataStore)
		list.dataStore = newData
	}
}
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("数组越界")
	}
	if list.TheSize != 0 {
		list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
		list.TheSize--
	}
	return nil
}
func (list *ArrayList) Size() int {
	return list.TheSize
}
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
