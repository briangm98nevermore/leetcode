package main

import "leetcode/treeutils"


func sortedArrayToBST(nums []int) *treeutils.TreeNode {
	return build(nums, 0, len(nums)-1)
}


func build(nums []int, izquierda, derecha int) *treeutils.TreeNode {
	
	if izquierda > derecha {
		return nil
	}
	
	medio := (izquierda + derecha) / 2	
	node := &treeutils.TreeNode{Val: nums[medio]}
		
	node.Left = build(nums, izquierda, medio-1)	
	node.Right = build(nums, medio+1, derecha)
	
	return node
}

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	treeutils.PrintTree(root, "", true)
}