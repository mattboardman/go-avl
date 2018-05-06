package avl

import (
	"reflect"
	"testing"
)

func TestTreeInsert(t *testing.T) {
	t.Run("Appending node greater than parent", func(t *testing.T) {
		parent := NewTree(5)
		child := NewTree(6)
		parent.Insert(child)

		if parent.Right.Value != child.Value {
			t.Errorf("Did not append child node [%d] on the right of root [%d]", child.Value, parent.Value)
		}
	})

	t.Run("Appending node less than parent", func(t *testing.T) {
		parent := NewTree(5)
		child := NewTree(4)
		parent.Insert(child)

		if parent.Left.Value != child.Value {
			t.Errorf("Did not append child node [%d] on the left of root [%d]", child.Value, parent.Value)
		}
	})

	t.Run("Right Height of tree after 1 insert", func(t *testing.T) {
		parent := NewTree(5)
		parent.Insert(NewTree(6))

		got := parent.RightHeight
		want := 1

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}
	})

	t.Run("Right Height of tree after 3 inserts", func(t *testing.T) {
		parent := NewTree(5)
		parent.Insert(NewTree(6))
		parent.Insert(NewTree(3))
		parent.Insert(NewTree(7))

		got := parent.RightHeight
		want := 2

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}
	})
}

func TestRotateRight(t *testing.T) {
	t.Run("Rotate Tree Right", func(t *testing.T) {
		parent := NewTree(5)
		parent.Insert(NewTree(3))
		parent.Insert(NewTree(1))

		parent = parent.RotateRight()

		got := parent.Value
		want := 3

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}

		if parent.RightHeight != 1 || parent.RightHeight != parent.LeftHeight {
			t.Errorf("Heights are mismatched: left[%d] != right[%d]", parent.LeftHeight, parent.RightHeight)
		}
	})

	t.Run("Rotate Tree Right With Subtrees", func(t *testing.T) {
		parent := NewTree(6)      // Root
		parent.Insert(NewTree(7)) // A
		parent.Insert(NewTree(4)) // Left
		parent.Insert(NewTree(5)) // B
		parent.Insert(NewTree(2)) // Left Left

		parent.Insert(NewTree(3)) // C
		parent.Insert(NewTree(1)) // D

		parent = parent.RotateRight()

		got := []int{7, 5, 3, 1} // A B C D
		want := []int{parent.Right.Right.Value,
			parent.Right.Left.Value,
			parent.Left.Right.Value,
			parent.Left.Left.Value}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d Want %d", got, want)
		}
	})
}

func TestRotateLeft(t *testing.T) {
	t.Run("Rotate Tree Left", func(t *testing.T) {
		parent := NewTree(5)
		parent.Insert(NewTree(6))
		parent.Insert(NewTree(7))

		parent = parent.RotateLeft()

		got := parent.Value
		want := 6

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}

		if parent.RightHeight != 1 || parent.RightHeight != parent.LeftHeight {
			t.Errorf("Heights are mismatched: left[%d] != right[%d]", parent.LeftHeight, parent.RightHeight)
		}
	})

	t.Run("Rotate Tree Left With Subtrees", func(t *testing.T) {
		parent := NewTree(4)      // Root
		parent.Insert(NewTree(3)) // A
		parent.Insert(NewTree(6)) // Right
		parent.Insert(NewTree(5)) // B
		parent.Insert(NewTree(8)) // Right Right
		parent.Insert(NewTree(7)) // C
		parent.Insert(NewTree(9)) // D

		parent = parent.RotateLeft()

		got := []int{3, 5, 7, 9} // A B C D
		want := []int{parent.Left.Left.Value,
			parent.Left.Right.Value,
			parent.Right.Left.Value,
			parent.Right.Right.Value}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d Want %d", got, want)
		}
	})
}

func TestRotateRightLeft(t *testing.T) {
	t.Run("Rotate Tree Left-Right", func(t *testing.T) {
		parent := NewTree(5)      // Root
		parent.Insert(NewTree(7)) // Right
		parent.Insert(NewTree(6)) // Right Left

		parent = parent.RotateRightLeft()

		got := parent.Value
		want := 6

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}
	})

	t.Run("Rotate Tree Left-Right With Subtrees", func(t *testing.T) {
		parent := NewTree(8)      // Root
		parent.Insert(NewTree(9)) // A
		parent.Insert(NewTree(4)) // Left
		parent.Insert(NewTree(1)) // B
		parent.Insert(NewTree(6)) // Left Right

		parent.Insert(NewTree(5)) // C
		parent.Insert(NewTree(7)) // D

		parent = parent.RotateLeftRight()

		got := []int{9, 1, 5, 7} // A B C D
		want := []int{parent.Right.Right.Value,
			parent.Left.Left.Value,
			parent.Left.Right.Value,
			parent.Right.Left.Value}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d Want %d", got, want)
		}
	})
}

func TestRotateLeftRight(t *testing.T) {
	t.Run("Rotate Tree Left-Right", func(t *testing.T) {
		parent := NewTree(5)      // Root
		parent.Insert(NewTree(3)) // Left
		parent.Insert(NewTree(4)) // Left Right

		parent = parent.RotateLeftRight()

		got := parent.Value
		want := 4

		if got != want {
			t.Errorf("Got %d Want %d", got, want)
		}
	})

	t.Run("Rotate Tree Left-Right With Subtrees", func(t *testing.T) {
		parent := NewTree(8)      // Root
		parent.Insert(NewTree(9)) // A
		parent.Insert(NewTree(4)) // Left
		parent.Insert(NewTree(1)) // B
		parent.Insert(NewTree(6)) // Left Right

		parent.Insert(NewTree(5)) // C
		parent.Insert(NewTree(7)) // D

		parent = parent.RotateLeftRight()

		got := []int{9, 1, 5, 7} // A B C D
		want := []int{parent.Right.Right.Value,
			parent.Left.Left.Value,
			parent.Left.Right.Value,
			parent.Right.Left.Value}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d Want %d", got, want)
		}
	})
}
