package collections

type Map[K comparable, V any] interface {
	ForEach(func(key K, value V))
}
