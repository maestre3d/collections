package lists

import "github.com/maestre3d/collections"

type List[T comparable] interface {
	collections.Collection[T]
	ListIterable[T]
	AddAllWithIndex(idx int, list collections.Collection[T]) bool
	Get(idx int) T
	Set(idx int, v T) T
	RemoveWithIndex(idx int) T
	IndexOf(v T) int
	LastIndexOf(v T) int
}
