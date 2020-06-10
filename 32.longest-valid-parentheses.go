/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (27.58%)
 * Likes:    3293
 * Dislikes: 132
 * Total Accepted:    276.5K
 * Total Submissions: 989.2K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 * 
 * Example 1:
 * 
 * 
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 * 
 * 
 */

// @lc code=start

// 利用栈实现
func longestValidParentheses(s string) int {
	l := list.New()
	l.PushBack(-1)
	max := 0
	for i, _ := range s {
		switch s[i] {
		case '(':
			{
				l.PushBack(i)
			}
		case ')':
			{
				l.Remove(l.Back())
				if l.Len() == 0 {
					l.PushBack(i)
				} else {
					j := l.Back().Value.(int)
					if max < i-j {
						max = i - j
					}
				}
			}
		}
	}
	return max
}

// 利用左右括号计数器实现
func longestValidParentheses(s string) int {
	left := 0
	right := 0
	max := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			left++
		case ')':
			right++
		}
		if right == left {
			if max < left+right {
				max = left + right
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			left++
		case ')':
			right++
		}
		if right == left {
			if max < left+right {
				max = left + right
			}
		} else if left > right {
			left, right = 0, 0
		}
	}

	return max
}
// @lc code=end

