/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (51.94%)
 * Likes:    3811
 * Dislikes: 561
 * Total Accepted:    941.2K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 * 
 * Example:
 * 
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
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
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = &ListNode{
					Val: l1.Val,
				}
				l1 = l1.Next
			} else {
				cur.Next = &ListNode{
					Val: l2.Val,
				}
				l2 = l2.Next
			}
		} else if l1 != nil {
			cur.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return head.Next
}
// @lc code=end

