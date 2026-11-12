package main

import (
	"fmt"
	"leetcode/internal/listutils"
)

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached
again by continuously following the next pointer. Internally, pos is used to denote the index
of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

Constraints:

    The number of the nodes in the list is in the range [0, 104].
    -105 <= Node.val <= 105
    pos is -1 or a valid index in the linked-list.


Follow up: Can you solve it using O(1) (i.e. constant) memory?m
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *listutils.ListNode) bool {

	if head == nil {

		return false
	}

	if head.Next == nil {
		return false
	}

	if head.Next != nil {

		if head.Next == head {
			return true
		}

		if head.Next.Next == head {
			return true
		}
	}

	if head.Next.Next == nil {
		return false
	}

	cabeza, cola := head, head

	for cola != nil && cola.Next != nil {

		cabeza = cabeza.Next
		cola = cola.Next.Next

		if cabeza == cola {
			return true
		}

	}

	return false
}

func main() {

	slice := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}

	simplelist := listutils.CreateList(slice)

	// head := &listutils.ListNode{Val: 1}
	// node1 := &listutils.ListNode{Val: 2}
	// node2 := &listutils.ListNode{Val: 3}
	// node3 := &listutils.ListNode{Val: 4}
	// node4 := &listutils.ListNode{Val: 5}
	// node5 := &listutils.ListNode{Val: 6}

	// head.Next = node1
	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5
	// node5.Next = node2

	fmt.Println(hasCycle(simplelist))
}
