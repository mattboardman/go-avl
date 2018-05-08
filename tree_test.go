package avl

import (
	"math/rand"
	"testing"
)

func TestTreeRebalance(t *testing.T) {
	test := Tree{}
	test.Insert(NewTree(5))

	for i := 0; i < 1000; i++ {
		test.Insert(NewTree(rand.Int()))
	}

	if test.Root.Balance > 1 || test.Root.Balance < -1 {
		t.Errorf("Tree unbalanced: %d", test.Root.Balance)
	}
}
