package collections

type Collection[T comparable] interface {
	Iterable[T]
	Len() int
	IsEmpty() bool
	Contains(v T) bool
	ToSlice() []T
	Add(v T) bool
	Remove(v T) bool
	ContainsAll(coll Collection[T]) bool
	AddAll(coll Collection[T]) bool
	RemoveAll(coll Collection[T]) bool
	RemoveIf(predicateFunc PredicateFunc[T]) bool
	RetainAll(coll Collection[T]) bool
	Clear()
}
