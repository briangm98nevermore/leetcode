package main

import (
	"fmt"
)

// ListNode es la definición de la lista enlazada (copiada del comentario de LeetCode)
type ListNode struct {
    Val  int
    Next *ListNode
}

// deleteDuplicates es la función que implementas (tal cual aparecerá en LeetCode)
func deleteDuplicates(head *ListNode) *ListNode {
   
	if head == nil {
		return nil
	}
	
	current := head

	for current.Next != nil{
			if current.Val == current.Next.Val {
			current.Next = current.Next.Next
			}else{current = current.Next}
		
	}

	return head
}

// --- Funciones auxiliares para pruebas ---
func buildList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    head := &ListNode{Val: nums[0]}
    tail := head
    for _, v := range nums[1:] {
        tail.Next = &ListNode{Val: v}
        tail = tail.Next
    }
    return head
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d ", head.Val)
        head = head.Next
    }
    fmt.Println()
}

func main() {
    // Prueba local
    lista := buildList([]int{1, 1, 2, 3, 3})
    fmt.Print("Original: ")
    printList(lista)

    lista = deleteDuplicates(lista)
    fmt.Print("Sin duplicados: ")
    printList(lista)

}

// x := 5
// p := &x
// fmt.Println(*p)

/*

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    current := head
    for current.Next != nil {
        if current.Val == current.Next.Val {
            // Saltamos el nodo duplicado
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    return head
}


*/