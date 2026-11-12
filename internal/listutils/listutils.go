package listutils

import "fmt"

// ListNode representa un nodo de una lista
type ListNode struct{
	Val int
	Next *ListNode
}

func CreateList(nums []int) *ListNode {

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

func PrintList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d -> ", head.Val)
        head = head.Next
    }
    fmt.Println("nil")
}