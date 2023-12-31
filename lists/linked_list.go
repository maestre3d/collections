package lists

import (
	"container/list"

	"github.com/maestre3d/collections"
)

type LinkedList[T comparable] struct {
	buf *list.List
}

var _ List[string] = LinkedList[string]{}

func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		buf: list.New(),
	}
}

func (l LinkedList[T]) Iterator() collections.Iterator[T] {
	return NewIterator[T](l)
}

func (l LinkedList[T]) Len() int {
	return l.buf.Len()
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.buf.Len() == 0
}

func (l LinkedList[T]) Contains(v T) bool {
	currNode := l.buf.Front()
	for currNode != nil {
		if currNode.Value.(T) == v {
			return true
		}

		currNode = currNode.Next()
	}

	return false
}

func (l LinkedList[T]) ToSlice() []T {
	iter := l.Iterator()
	buf := make([]T, 0, l.Len())
	for iter.HasNext() {
		buf = append(buf, iter.Next())
	}
	return buf
}

func (l LinkedList[T]) Add(v T) bool {
	return l.buf.PushBack(v) != nil
}

func (l LinkedList[T]) Remove(v T) bool {
	idx := l.IndexOf(v)
	elem := l.get(idx)
	item := l.buf.Remove(elem).(T)
	var zeroVal T
	return item != zeroVal
}

func (l LinkedList[T]) ContainsAll(coll collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func newListFromCollection[T comparable](coll collections.Collection[T]) *list.List {
	iter := coll.Iterator()
	buf := list.New()
	for iter.HasNext() {
		buf.PushBack(iter.Next())
	}

	return buf
}

func (l LinkedList[T]) AddAll(coll collections.Collection[T]) bool {
	buf := newListFromCollection(coll)
	l.buf.PushBackList(buf)
	return true
}

func (l LinkedList[T]) RemoveAll(coll collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList[T]) RemoveIf(predicateFunc collections.PredicateFunc[T]) bool {
	currNode := l.buf.Front()
	for currNode != nil {
		val := currNode.Value.(T)
		if predicateFunc(val) {
			l.Remove(val)
			return true
		}
		currNode = currNode.Next()
	}

	return false
}

func (l LinkedList[T]) RetainAll(coll collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList[T]) Clear() {
	l.buf.Init()
}

func (l LinkedList[T]) ListIterator() Iterator[T] {
	return NewIterator[T](l)
}

func (l LinkedList[T]) AddAllWithIndex(idx int, list collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList[T]) get(idx int) *list.Element {
	currNode := l.buf.Front()
	currIdx := 0
	for currNode != nil {
		if currIdx == idx {
			return currNode
		}
		currNode = currNode.Next()
		currIdx++
	}

	return nil
}

func (l LinkedList[T]) Get(idx int) T {
	elem := l.get(idx)
	if elem != nil {
		return elem.Value.(T)
	}

	var zeroVal T
	return zeroVal
}

func (l LinkedList[T]) Set(idx int, v T) T {
	elem := l.get(idx)
	elem.Value = v
	return v
}

func (l LinkedList[T]) RemoveWithIndex(idx int) T {
	elem := l.get(idx)
	l.buf.Remove(elem)
	return elem.Value.(T)
}

func (l LinkedList[T]) IndexOf(v T) int {
	currNode := l.buf.Front()
	currIdx := 0
	for currNode != nil {
		val := currNode.Value.(T)
		if val == v {
			return currIdx
		}
		currNode = currNode.Next()
		currIdx++
	}

	return -1
}

func (l LinkedList[T]) LastIndexOf(v T) int {
	currNode := l.buf.Front()
	lastSeenIdx := -1
	currIdx := 0
	for currNode != nil {
		val := currNode.Value.(T)
		if val == v {
			lastSeenIdx = currIdx
		}
		currNode = currNode.Next()
		currIdx++
	}

	return lastSeenIdx
}
