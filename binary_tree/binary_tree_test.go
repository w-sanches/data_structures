package binarytree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	binarytree "github.com/w-sanches/data_structures/binary_tree"
)

func TestInsert(t *testing.T) {
	tree := binarytree.New(50)

	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(0)
	tree.Insert(80)
	tree.Insert(100)
	tree.Insert(70)

	assert.Equal(t, tree.Value, 50)
	assert.Equal(t, tree.Left.Value, 20)
	assert.Equal(t, tree.Left.Left.Value, 0)
	assert.Equal(t, tree.Left.Right.Value, 30)
	assert.Equal(t, tree.Right.Value, 80)
	assert.Equal(t, tree.Right.Right.Value, 100)
	assert.Equal(t, tree.Right.Left.Value, 70)
}
