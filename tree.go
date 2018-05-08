package avl

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(newNode *Node) {
	if t.Root == nil {
		t.Root = newNode
		return
	}
	t.Root.Insert(newNode)
	if t.Root.Balance > 1 || t.Root.Balance < -1 {
		t.Rebalance()
	}
}

func (t *Tree) Rebalance() {
	fake := &Node{Left: t.Root}
	fake.Rebalance(t.Root)
	t.Root = fake.Left
}
