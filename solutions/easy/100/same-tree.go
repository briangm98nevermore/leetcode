/**

Given the roots of two binary trees p and q, 
write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, 
and the nodes have the same value.

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"
import "leetcode/internal/treeutils"

func isSameTree(p *treeutils.TreeNode, q *treeutils.TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main()  {

	p := treeutils.CreateTree([]int{1,2,1})
	q := treeutils.CreateTree([]int{1,1,2})

	treeutils.PrintTree(p, "", false)
	treeutils.PrintTree(q, "", false)

	fmt.Println(isSameTree(p,q))

}