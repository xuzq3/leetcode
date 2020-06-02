/*
 * Given a linked list, swap every two adjacent nodes and return its head.
 * You may not modify the values in the list's nodes, only nodes itself may be changed.
 * 
 * Example:
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var n *ListNode = &ListNode{
		Next: head,
	}
	var pre *ListNode = n
	var tmp *ListNode = nil

	for pre.Next != nil {
		tmp = pre.Next
		if tmp.Next == nil {
			break
		}

		pre.Next = tmp.Next
		tmp.Next = tmp.Next.Next
		pre.Next.Next = tmp

		pre = tmp
	}

	return n.Next
}
