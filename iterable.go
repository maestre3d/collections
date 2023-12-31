package collections

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type CollectionIterator[T comparable] struct {
	buf Collection[T]
}

var _ Iterator[string] = &CollectionIterator[string]{}

func (c *CollectionIterator[T]) HasNext() bool {
	panic("impl me")
}

func (c *CollectionIterator[T]) Next() T {
	//TODO implement me
	panic("implement me")
}
