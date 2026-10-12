package treeutils

import "fmt"

// TreeNode representa un nodo de árbol binario.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateTree construye un árbol binario a partir de un slice.
// -1 significa nodo nulo.
func CreateTree(nums []int) *TreeNode {
	
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}

	i := 1
	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) && nums[i] != -1 {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != -1 {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// PrintTree imprime un árbol de forma visual.
func PrintTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Print("└── ")
	} else {
		fmt.Print("├── ")
	}
	fmt.Println(root.Val)

	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}

	PrintTree(root.Left, newPrefix, true)
	PrintTree(root.Right, newPrefix, false)
}