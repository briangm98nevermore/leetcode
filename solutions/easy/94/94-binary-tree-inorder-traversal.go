/**
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]



 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

package main

import (
	"fmt"
    "leetcode/treeutils"
)

func inorderTraversal(root *treeutils.TreeNode) []int {

    if root == nil {
        return nil
    }
    
    result := []int{}
    stack := []*treeutils.TreeNode{}
    current := root

    for current != nil || len(stack) > 0 {
        // Ir lo más a la izquierda posible, apilando los nodos
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        // Sacar el último nodo apilado
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // Procesar el nodo
        result = append(result, current.Val)
        // Ahora ir al hijo derecho
        current = current.Right
    }
    return result
}
func main()  {

	raiz := treeutils.CreateTree([]int{1,2,3,4,5,-1,8,-1,-1,6,7,9})
	//printTree(raiz, "", false)
    result:=inorderTraversal(raiz)
    fmt.Println(result)


}