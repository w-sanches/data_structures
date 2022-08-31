package singlelinkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	singlelinkedlist "github.com/w-sanches/data_structures/single_linked_list"
)

func TestNew(t *testing.T) {
	list := singlelinkedlist.NewSingleLinked[string]()
	assert.Nil(t, list.Root)
	otherList := singlelinkedlist.NewSingleLinked[int]()
	assert.Nil(t, otherList.Root)
}

func TestAdd(t *testing.T) {
	list := singlelinkedlist.NewSingleLinked[string]()
	list.Add("two")
	assert.Equal(t, list.Root.Value, "two")

	list.Add("one")
	assert.Equal(t, list.Root.Value, "one")
}

func TestAppend(t *testing.T) {
	list := singlelinkedlist.NewSingleLinked[string]()
	list.Append("one")
	list.Append("two")
	first := list.Root
	second := first.Next

	assert.Equal(t, first.Value, "one")
	assert.Equal(t, second.Value, "two")
}

func TestFirst(t *testing.T) {
	list := singlelinkedlist.NewSingleLinked[string]()
	list.Append("one")
	list.Append("two")
	list.Append("three")

	assert.Equal(t, list.First(), "one")
}

func TestLength(t *testing.T) {
	t.Run("Is zero if list is empty", func(t *testing.T) {
		list := singlelinkedlist.NewSingleLinked[string]()
		assert.Equal(t, list.Length(), 0)
	})

	list := singlelinkedlist.NewSingleLinked[string]()
	list.Append("two")
	list.Add("one")
	list.Append("three")
	assert.Equal(t, list.Length(), 3)
}
func TestLast(t *testing.T) {
	list := singlelinkedlist.NewSingleLinked[string]()
	list.Append("one")
	list.Append("two")
	list.Append("three")

	assert.Equal(t, list.Last(), "three")
}

func TestRemove(t *testing.T) {
	t.Run("Does nothing if there are no elements to remove", func(t *testing.T) {
		list := singlelinkedlist.NewSingleLinked[string]()
		list.Append("one")
		list.Append("two")
		list.Remove("three")

		assert.NotNil(t, list.Root.Next)
		assert.Equal(t, list.Root.Next.Value, "two")
	})
	t.Run("Removes elements", func(t *testing.T) {
		list := singlelinkedlist.NewSingleLinked[string]()
		list.Append("one")
		list.Append("two")
		list.Remove("two")

		assert.Nil(t, list.Root.Next)
	})
	t.Run("Removes multiple elements of same value", func(t *testing.T) {
		list := singlelinkedlist.NewSingleLinked[string]()
		list.Append("one")
		list.Append("two")
		list.Append("two")
		list.Remove("two")

		assert.Nil(t, list.Root.Next)
	})
	t.Run("Removes elements when on root", func(t *testing.T) {
		list := singlelinkedlist.NewSingleLinked[string]()
		list.Append("one")
		list.Append("two")
		list.Remove("one")

		if assert.NotNil(t, list.Root) {
			assert.Equal(t, list.Root.Value, "two")
		}
	})
}
