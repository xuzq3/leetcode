/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (40.24%)
 * Likes:    2012
 * Dislikes: 354
 * Total Accepted:    258K
 * Total Submissions: 629.6K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 * 
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes in the end should remain as it is.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * Given this linked list: 1->2->3->4->5
 * 
 * For k = 2, you should return: 2->1->4->3->5
 * 
 * For k = 3, you should return: 3->2->1->4->5
 * 
 * Note:
 * 
 * 
 * Only constant extra memory is allowed.
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func reverseKGroup(head *ListNode, k int) *ListNode {
	var result *ListNode = &ListNode{
		Next: head,
	}
	var pre *ListNode = result	// 短链前一个node
	var end *ListNode = head	// 短链最后一个node
	var cur *ListNode = head	// 当前node
	var next *ListNode			// 下一个node
	i := 0
	for cur != nil {
		// 短链就地反转
		next = cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		end.Next = next
		cur = next

		i++
		if i == k {
			// 按长度 k 进行分段
			for i > 0 {
				pre = pre.Next
				i--
			}
			end = pre.Next
		}
	}
	cur = pre.Next
	end = cur
	for cur != nil {
		// 最后长度不足k部分再次反转到正常顺序
		next = cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		end.Next = next
		cur = next
	}
	return result.Next
}
// @lc code=end

