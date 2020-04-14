/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (28.94%)
 * Likes:    5943
 * Dislikes: 487
 * Total Accepted:    852.7K
 * Total Submissions: 2.9M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 * 
 * Example 1:
 * 
 * 
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "cbbd"
 * Output: "bb"
 * 
 * 
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	sub := ""
	tmp := ""
	for i := 0; i < n; i++ {
		left, right := i, i
		for right < n - 1 && s[right] == s[right+1] {
			right++
		}
		i = right
		tmp = expandAroundCenter(s, left, right)
		if len(tmp) > len(sub) {
			sub = tmp
		}
	}
	return sub
}

func longestPalindrome_ori(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	sub := ""
	tmp := ""
	for i := 0; i < n; i++ {
		tmp = expandAroundCenter(s, i, i)
		if len(tmp) > len(sub) {
			sub = tmp
		}
		tmp = expandAroundCenter(s, i, i+1)
		if len(tmp) > len(sub) {
			sub = tmp
		}
	}
	return sub
}

func expandAroundCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
// @lc code=end

