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

// PrintTree imprime el árbol rotado 90° con la raíz a la izquierda.
// El hijo derecho aparece arriba y el izquierdo abajo.
// Los hijos únicos muestran un slot vacío (·) del lado que falta.
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("(árbol vacío)")
		return
	}
	fmt.Println()
	fmt.Println(root.Val)
	printChildren(root, "")
	fmt.Println()
}

// printChildren decide cómo dibujar los hijos de un nodo:
//   - ambos: hijo derecho arriba (┌──), izquierdo abajo (└──)
//   - solo derecho: hijo arriba, slot vacío abajo (└── ·)
//   - solo izquierdo: slot vacío arriba (┌── ·), hijo abajo
//   - ninguno: nada
func printChildren(node *TreeNode, prefix string) {
	hasLeft := node.Left != nil
	hasRight := node.Right != nil

	switch {
	case hasLeft && hasRight:
		printRotated(node.Right, prefix, false)
		printRotated(node.Left, prefix, true)
	case hasRight:
		printRotated(node.Right, prefix, false)
		fmt.Println(prefix + "└── ·")
	case hasLeft:
		fmt.Println(prefix + "┌── ·")
		printRotated(node.Left, prefix, true)
	}
	// sin hijos: no imprimir nada
}

// printRotated imprime un nodo y delega sus hijos a printChildren.
func printRotated(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	connector := "┌── "
	if isLeft {
		connector = "└── "
	}
	fmt.Println(prefix + connector + fmt.Sprint(node.Val))

	newPrefix := prefix
	if isLeft {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	printChildren(node, newPrefix)
}