package lists

import "github.com/maestre3d/collections"

type ListSlice[T comparable] struct {
	buf []T
}

var _ List[string] = &ListSlice[string]{}

func NewListSlice[T comparable](cap int) *ListSlice[T] {
	if cap < 0 {
		cap = 0
	}
	return &ListSlice[T]{
		buf: make([]T, cap),
	}
}

func (l *ListSlice[T]) Iterator() collections.Iterator[T] {
	return NewIterator[T](l)
}

func (l *ListSlice[T]) Len() int {
	return len(l.buf)
}

func (l *ListSlice[T]) IsEmpty() bool {
	return len(l.buf) == 0
}

func (l *ListSlice[T]) Contains(v T) bool {
	for _, item := range l.buf {
		if v == item {
			return true
		}
	}

	return false
}

func (l *ListSlice[T]) ToSlice() []T {
	return l.buf
}

func (l *ListSlice[T]) Add(v T) bool {
	l.buf = append(l.buf, v)
	return true
}

func (l *ListSlice[T]) Remove(v T) bool {
	idx := l.IndexOf(v)
	var zeroVal T
	return zeroVal != l.RemoveWithIndex(idx)
}

func (l *ListSlice[T]) ContainsAll(coll collections.Collection[T]) bool {
	panic("imp")
}

func (l *ListSlice[T]) AddAll(coll collections.Collection[T]) bool {
	l.buf = append(l.buf, coll.ToSlice()...)
	return true
}

func (l *ListSlice[T]) RemoveAll(coll collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *ListSlice[T]) RemoveIf(predicateFunc collections.PredicateFunc[T]) bool {
	wasMod := false
	var zeroVal T
	for i, item := range l.buf {
		if ok := predicateFunc(item); ok {
			wasMod = zeroVal != l.RemoveWithIndex(i)
		}
	}

	return wasMod
}

func (l *ListSlice[T]) RetainAll(coll collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *ListSlice[T]) Clear() {
	l.buf = l.buf[:0]
}

func (l *ListSlice[T]) ListIterator() Iterator[T] {
	return NewIterator[T](l)
}

func (l *ListSlice[T]) AddAllWithIndex(idx int, list collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *ListSlice[T]) Get(idx int) T {
	if idx >= len(l.buf) {
		var zeroVal T
		return zeroVal
	}
	return l.buf[idx]
}

func (l *ListSlice[T]) Set(idx int, v T) T {
	if idx >= len(l.buf) {
		var zeroVal T
		return zeroVal
	}

	l.buf[idx] = v
	return v
}

func (l *ListSlice[T]) RemoveWithIndex(idx int) T {
	if idx < len(l.buf)-1 {
		copy(l.buf[idx:], l.buf[idx+1:])
	}

	var zeroVal T
	l.buf[len(l.buf)-1] = zeroVal
	l.buf = l.buf[:len(l.buf)-1]
	return zeroVal
}

func (l *ListSlice[T]) IndexOf(v T) int {
	for i, item := range l.buf {
		if item == v {
			return i
		}
	}

	return -1
}

func (l *ListSlice[T]) LastIndexOf(v T) int {
	lastSeenItemIdx := -1
	for i, item := range l.buf {
		if item == v {
			lastSeenItemIdx = i
		}
	}

	return lastSeenItemIdx
}
