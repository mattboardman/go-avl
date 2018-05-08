package avl

type Node struct {
	Left    *Node
	Right   *Node
	Balance int
	Value   int
}

func NewTree(value int) *Node {
	return &Node{Value: value}
}

func (root *Node) Insert(newNode *Node) bool {
	switch {
	case newNode.Value == root.Value:
		return false
	case newNode.Value < root.Value:
		if root.Left == nil {
			root.Left = newNode
			if root.Right == nil {
				root.Balance = -1
			} else {
				root.Balance = 0
			}
		} else {
			if root.Left.Insert(newNode) {
				if root.Left.Balance < -1 || root.Left.Balance > 1 {
					root.Rebalance(root.Left)
				} else {
					root.Balance--
				}
			}
		}
	case newNode.Value > root.Value:
		if root.Right == nil {
			root.Right = newNode
			if root.Left == nil {
				root.Balance = 1
			} else {
				root.Balance = 0
			}
		} else {
			if root.Right.Insert(newNode) {
				if root.Right.Balance < -1 || root.Right.Balance > 1 {
					root.Rebalance(root.Right)
				} else {
					root.Balance++
				}
			}
		}
	}

	if root.Balance != 0 {
		return true
	}
	return false
}

func (root *Node) Rebalance(child *Node) {
	switch {
	case child.Balance == -2 && child.Left.Balance == -1:
		root.RotateRight(child)
	case child.Balance == 2 && child.Right.Balance == 1:
		root.RotateLeft(child)
	case child.Balance == -2 && child.Left.Balance == 1:
		root.RotateLeftRight(child)
	case child.Balance == 2 && child.Right.Balance == -1:
		root.RotateRightLeft(child)
	}
}

func (root *Node) RotateRight(child *Node) {
	center := child.Left
	child.Left = center.Right
	center.Right = child

	if child == root.Left {
		root.Left = center
	} else {
		root.Right = center
	}
	center.Balance = 0
	child.Balance = 0
}

func (root *Node) RotateLeft(child *Node) {
	center := child.Right
	child.Right = center.Left
	center.Left = child

	if child == root.Left {
		root.Left = center
	} else {
		root.Right = center
	}
	center.Balance = 0
	child.Balance = 0
}

func (root *Node) RotateLeftRight(child *Node) {
	child.Left.Right.Balance = -1
	child.RotateLeft(child.Left)
	child.Left.Balance = -1
	root.RotateRight(child)
}

func (root *Node) RotateRightLeft(child *Node) {
	child.Right.Left.Balance = 1
	child.RotateRight(child.Right)
	child.Right.Balance = 1
	root.RotateLeft(child)
}
