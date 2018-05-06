package avl

type Tree struct {
	Left        *Tree
	Right       *Tree
	RightHeight int
	LeftHeight  int
	Value       int
}

func NewTree(value int) *Tree {
	return &Tree{Value: value}
}

func (root *Tree) Insert(newNode *Tree) {
	if newNode.Value > root.Value {
		if root.Right == nil {
			root.Right = newNode
		} else {
			root.Right.Insert(newNode)
		}
		root.RightHeight++
	} else if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			root.Left.Insert(newNode)
		}
		root.LeftHeight++
	}
}

func (root *Tree) RotateRight() *Tree {
	center := root.Left
	center.RightHeight = root.RightHeight + 1
	root.Left = center.Right
	center.Right = root

	return center
}

func (root *Tree) RotateLeft() *Tree {
	center := root.Right
	center.LeftHeight = root.LeftHeight + 1
	root.Right = center.Left
	center.Left = root

	return center
}

func (root *Tree) RotateLeftRight() *Tree {
	/* 	center := root.Left.Right
	   	center.LeftHeight = root.LeftHeight + 1
	   	root.Left.Right = center.Left
	   	center.Left = root.Left
	   	root.Left = center */

	root.Left = root.Left.RotateLeft()
	return root.RotateRight()
}

func (root *Tree) RotateRightLeft() *Tree {
	/* 	center := root.Right.Left
	   	center.RightHeight = root.RightHeight + 1
	   	root.Right.Left = center.Right
	   	center.Right = root.Right
	   	root.Right = center */

	root.Right = root.Right.RotateRight()

	return root.RotateLeft()
}
