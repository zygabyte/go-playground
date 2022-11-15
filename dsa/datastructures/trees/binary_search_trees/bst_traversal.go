package main

import "fmt"

type BST struct {
	value       int
	left, right *BST
}

func main() {
	bst := &BST{
		value: 10,
		left: &BST{
			value: 5,
			left: &BST{
				value: 2,
				left: &BST{
					value: 1,
				},
			}, right: &BST{
				value: 5,
			},
		}, right: &BST{
			value: 15,
			right: &BST{
				value: 22,
			},
		},
	}

	array := []int{}
	in := bst.inorderTravesal(array)
	pre := bst.preorderTravesal(array)
	post := bst.postorderTravesal(array)

	fmt.Println("in -> ", in)
	fmt.Println("pre -> ", pre)
	fmt.Println("post -> ", post)
}

func (b *BST) inorderTravesal(array []int) []int {
	if b.left != nil {
		array = b.left.inorderTravesal(array)
	}
	array = append(array, b.value)
	if b.right != nil {
		array = b.right.inorderTravesal(array)
	}

	return array
}

func (b *BST) preorderTravesal(array []int) []int {
	array = append(array, b.value)
	if b.left != nil {
		array = b.left.preorderTravesal(array)
	}
	if b.right != nil {
		array = b.right.preorderTravesal(array)
	}

	return array
}

func (b *BST) postorderTravesal(array []int) []int {
	if b.left != nil {
		array = b.left.postorderTravesal(array)
	}
	if b.right != nil {
		array = b.right.postorderTravesal(array)
	}
	array = append(array, b.value)

	return array
}
