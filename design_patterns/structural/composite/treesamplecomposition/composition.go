package treesamplecomposition

type Tree[T any] struct {
	LeafValue T
	Left      *Tree[T]
	Right     *Tree[T]
}

func NewTree[T any](leafValue T) Tree[T] {
	return Tree[T]{leafValue, nil, nil}
}
