package main

import (
	"fmt"
	"leetcode/internal/treeutils"
)

func isSymmetric(root *treeutils.TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil || root.Right == nil {
		return false
	} else if root.Right.Val != root.Left.Val {
		return false
	}

	sons := []*treeutils.TreeNode{}
	sons = append(sons, root.Left, root.Right)

	for len(sons) > 0 {

		node1, node2 := sons[0],sons[1]
		sons = sons[2:]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		sons = append(sons, node1.Left, node2.Right, node1.Right,node2.Left)
	}

	return true
}

func main() {

	root := treeutils.CreateTree([]int{1, 2, 2, 3, -1, -1, 3})
	//root := treeutils.CreateTree([]int{1,2,2,-1,3,-1,3})
	treeutils.PrintTree(root, "", false)

	fmt.Println(isSymmetric(root))

}
