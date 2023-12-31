package lists

import "github.com/maestre3d/collections"

type Iterator[T any] interface {
	collections.Iterator[T]
	HasPrevious() bool
	Previous() T
	NextIndex() int
	PreviousIndex() int
}

type ListIterable[T any] interface {
	ListIterator() Iterator[T]
}

type ListIterator[T comparable] struct {
	forwardCurrIdx  int
	backwardCurrIdx int
	buf             List[T]
}

var _ Iterator[string] = &ListIterator[string]{}

func NewIterator[T comparable](src List[T]) *ListIterator[T] {
	return &ListIterator[T]{
		forwardCurrIdx:  -1,
		backwardCurrIdx: src.Len(),
		buf:             src,
	}
}

func (i *ListIterator[T]) HasNext() bool {
	return i.forwardCurrIdx < i.buf.Len()-1
}

func (i *ListIterator[T]) Next() T {
	i.forwardCurrIdx++
	item := i.buf.Get(i.forwardCurrIdx)
	return item
}

func (i *ListIterator[T]) HasPrevious() bool {
	return i.backwardCurrIdx > 0 && i.backwardCurrIdx <= i.buf.Len()
}

func (i *ListIterator[T]) Previous() T {
	i.backwardCurrIdx--
	return i.buf.Get(i.backwardCurrIdx)
}

func (i *ListIterator[T]) NextIndex() int {
	i.forwardCurrIdx++
	return i.forwardCurrIdx
}

func (i *ListIterator[T]) PreviousIndex() int {
	i.backwardCurrIdx--
	return i.backwardCurrIdx
}
