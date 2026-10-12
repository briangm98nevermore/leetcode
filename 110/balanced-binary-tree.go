package main

import (
	"fmt"
	"leetcode/treeutils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

	Given a binary tree, determine if it is height-balanced.

	Example 1:
	Input: root = [3,9,20,null,null,15,7]
	Output: true

	Example 2:
	Input: root = [1,2,2,3,3,null,null,4,4]
	Output: false

	Example 3:

	Input: root = []
	Output: true

	Constraints:

		The number of nodes in the tree is in the range [0, 5000].
		-104 <= Node.val <= 104

*/
func isBalanced(root *treeutils.TreeNode) bool {

   if root == nil {
	 return true
   }

   balance, _ := checkdepth(root)

   return balance
}

func checkdepth(node *treeutils.TreeNode) (bool,int)  {
	
	if node == nil {
		return true, 0
	}

	izqBal, altIzq := checkdepth(node.Left)
	derBal, altDer := checkdepth(node.Right)

	diference := altIzq - altDer
	if diference < 0 {
		diference = -diference
	}

	balance := izqBal && derBal && diference <= 1

	height := 1 + maximum(altIzq,altDer)

	return balance,height
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	
	root := treeutils.CreateTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(checkdepth(root))
	treeutils.PrintTree(root, "", false)
}