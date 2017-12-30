package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func TraversalHelper(result *[]int, root *TreeNode) {
	if root.Left != nil {
		TraversalHelper(result, root.Left)
	}
	*result = append(*result, root.Val)
	if root.Right != nil {
		TraversalHelper(result, root.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	TraversalHelper(&result, root)
	return result
}

func main() {
	//var head *TreeNode = nil
	head := new(TreeNode)
	head.Val = 1
	head.Right = new(TreeNode)
	head.Right.Val = 2
	head.Right.Left = new(TreeNode)
	head.Right.Left.Val = 3

	fmt.Println(inorderTraversal(head))
}
