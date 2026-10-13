package main

import (
	"fmt"
	"leetcode/internal/treeutils"
)

func maxDepth(root *treeutils.TreeNode) int {

	if root == nil {
		return 0
	}

	sons := []*treeutils.TreeNode{root}
	levels := 0

	for len(sons) > 0 {

		levelSize := len(sons)
		levels++

		for i := 0; i < levelSize; i++ {
			node := sons[0]
			sons = sons[1:]

			if node.Left != nil {
				sons = append(sons, node.Left)
			}
			if node.Right != nil {
				sons = append(sons, node.Right)
			}
		}
	}

	return levels
	
}

func main()  {

	root := treeutils.CreateTree([]int{3,9,20,-1,-1,15,7})
	treeutils.PrintTree(root, "", true)

	fmt.Println(maxDepth(root))
}