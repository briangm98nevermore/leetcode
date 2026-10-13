package main

import (
	"fmt"
	"leetcode/internal/treeutils"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }


   Given a binary tree, find its minimum depth.
   The minimum depth is the number of nodes along the shortest path from
   the root node down to the nearest leaf node.

   Note: A leaf is a node with no children.

   Example 1:
   Input: root = [3,9,20,null,null,15,7]
   Output: 2

   Example 2:
   Input: root = [2,null,3,null,4,null,5,null,6]
   Output: 5

   Constraints:

       The number of nodes in the tree is in the range [0, 105].
       -1000 <= Node.val <= 1000


*/
func minDepth(root *treeutils.TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    if root.Left !=nil && root.Right == nil{
      return  1 + minDepth(root.Left)
    }

    if root.Right !=nil && root.Left == nil{
      return  1 + minDepth(root.Right)
    }

    return 1 + min(minDepth(root.Left),minDepth(root.Right))

}

func main()  {
	root := treeutils.CreateTree([]int{3,9,20,-1,-1,15,7})

	fmt.Println(minDepth(root))

	treeutils.PrintTree(root)
}
