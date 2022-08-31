package singlelinkedlist

import "golang.org/x/exp/constraints"

type SingleLinkedList[T constraints.Ordered] struct {
	Root *SingleLinkedNode[T]
}

type SingleLinkedNode[T constraints.Ordered] struct {
	Value T
	Next  *SingleLinkedNode[T]
}

func NewSingleLinked[T constraints.Ordered]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (list *SingleLinkedList[T]) Add(value T) {
	node := &SingleLinkedNode[T]{Value: value, Next: list.Root}
	list.Root = node
}

func (list *SingleLinkedList[T]) Append(value T) {
	newNode := &SingleLinkedNode[T]{Value: value}
	if list.Root == nil {
		list.Root = newNode
		return
	}
	node := list.Root
	for {
		if node.Next == nil {
			node.Next = newNode
			break
		}
		node = node.Next
	}
}

func (list *SingleLinkedList[T]) First() T {
	return list.Root.Value
}

func (list *SingleLinkedList[T]) Last() T {
	node := list.Root
	for {
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return node.Value
}

func (list *SingleLinkedList[T]) Length() int {
	if list.Root == nil {
		return 0
	}
	len := 0
	node := list.Root
	for {
		if node == nil {
			return len
		}
		len++
		node = node.Next
	}
}

func (list *SingleLinkedList[T]) Remove(value T) {
	for {
		if list.Root == nil {
			return
		}
		if list.Root.Next == nil || list.Root.Value != value {
			break
		}
		list.Root = list.Root.Next
	}
	node := list.Root
	for {
		if node == nil || node.Next == nil {
			break
		}
		if node.Next.Value == value {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
}
