package avl

import (
	"math/rand"
	"testing"
)

func TestNodeInsert(t *testing.T) {
	t.Run("Insert item == parent", func(t *testing.T) {
		parent := NewTree(5)

		if parent.Insert(NewTree(5)) {
			t.Error("Inserted == value")
		}
	})

	t.Run("Insert item > parent", func(t *testing.T) {
		parent := NewTree(5)
		parent.Insert(NewTree(6))

		got := parent.Right.Value
		want := 6

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}
	})

	t.Run("Insert item <  parent", func(t *testing.T) {
		parent := NewTree(5)
		parent.Insert(NewTree(4))

		got := parent.Left.Value
		want := 4

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}
	})
}

func TestNodeRebalance(t *testing.T) {
	t.Run("Test 100 items", func(t *testing.T) {
		parent := NewTree(10)
		for i := 0; i < 50; i++ {
			parent.Insert(NewTree(rand.Intn(50)))
		}

		if parent.Right.Balance > 1 || parent.Right.Balance < -1 {
			t.Errorf("Unbalanced: %d, Value: %d", parent.Right.Balance, parent.Right.Value)
		}

		if parent.Left.Balance > 1 || parent.Left.Balance < -1 {
			t.Errorf("Unbalanced: %d, Value: %d", parent.Left.Balance, parent.Left.Value)
		}
	})
}

func TestRotateRight(t *testing.T) {
	t.Run("Rotate Tree Right", func(t *testing.T) {
		parent := NewTree(5)
		parent.Left = NewTree(3)
		parent.Left.Left = NewTree(1)

		parent.Balance = -2
		parent.Left.Balance = -1
		parent.RotateRight(parent.Left)

		got := parent.Left.Value
		want := 1

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}

		/* 		if parent.Balance != 0 {
			t.Errorf("Unbalanced: %d", parent.Balance)
		} */
	})
}

func TestRotateLeft(t *testing.T) {
	t.Run("Rotate Tree Left", func(t *testing.T) {
		parent := NewTree(5)
		parent.Right = NewTree(6)
		parent.Right.Right = NewTree(7)

		parent.Balance = 2
		parent.Right.Balance = 1
		parent.RotateLeft(parent.Right)

		got := parent.Right.Value
		want := 7

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}

		/* 		if parent.Balance != 0 {
			t.Errorf("Unbalanced: %d", parent.Balance)
		} */
	})
}

func TestRotateRightLeft(t *testing.T) {
	t.Run("Rotate Tree Right-Left", func(t *testing.T) {
		parent := NewTree(5)      // Root
		parent.Right = NewTree(7) // Right
		parent.Right.Right = NewTree(9)
		parent.Right.Right.Left = NewTree(8) // Right Left

		parent.Balance = 2
		parent.Right.Balance = -1
		parent.RotateRightLeft(parent.Right)

		got := parent.Right.Value
		want := 8

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}

		/* 		if parent.Balance != 0 {
			t.Errorf("Unbalanced: %d", parent.Balance)
		} */
	})
}

func TestRotateLeftRight(t *testing.T) {
	t.Run("Rotate Tree Left-Right", func(t *testing.T) {
		parent := NewTree(6)          // Root
		parent.Left = NewTree(5)      // Left
		parent.Left.Left = NewTree(3) // Left Right
		parent.Left.Left.Right = NewTree(4)

		parent.Balance = -2
		parent.Left.Balance = 1
		parent.RotateLeftRight(parent.Left)

		got := parent.Left.Value
		want := 4

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}

		/* 		if parent.Balance != 0 {
			t.Errorf("Unbalanced: %d", parent.Balance)
		} */
	})
}
