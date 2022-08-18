package main

import "fmt"

type BstTree struct {
	val         int
	left, right *BstTree
}

func (bst *BstTree) Insert(val int) *BstTree {
	if bst == nil {
		return &BstTree{val: val}
	}
	if val < bst.val {
		bst.left = bst.left.Insert(val)
	}
	if val > bst.val {
		bst.right = bst.right.Insert(val)
	}

	return bst
}

func (bst *BstTree) Contains(val int) bool {
	switch {
	case bst == nil:
		return false
	case val < bst.val:
		return bst.left.Contains(val)
	case val > bst.val:
		return bst.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var bst *BstTree
	bst = bst.Insert(5)
	bst = bst.Insert(12)
	bst = bst.Insert(53)
	bst = bst.Insert(555)
	bst = bst.Insert(15)
	bst = bst.Insert(10)

	fmt.Println(bst.Contains(1))
	fmt.Println(bst.Contains(10))
	fmt.Println(bst.Contains(12))

}
