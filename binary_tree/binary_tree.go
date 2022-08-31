package binarytree

type BinaryTree[T int] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func New[T int](value T) *BinaryTree[T] {
	return &BinaryTree[T]{Value: value}
}

func (tree *BinaryTree[T]) Insert(value T) {
	if value <= tree.Value {
		if tree.Left == nil {
			tree.Left = New(value)
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = New(value)
		} else {
			tree.Right.Insert(value)
		}
	}
}
