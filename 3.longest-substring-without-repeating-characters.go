/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.77%)
 * Likes:    8274
 * Dislikes: 501
 * Total Accepted:    1.4M
 * Total Submissions: 4.7M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "abcabcbb"
 * Output: 3 
 * Explanation: The answer is "abc", with the length of 3. 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3. 
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	sub := ""
	maxLen := 0
	for _, c := range s {
		for j, subc := range sub {
			if c == subc {
				sub = sub[j+1:]
				break
			}
		}
		sub = append(sub, c)
		if len(sub) > maxLen {
			maxLen = len(sub)
		}
	}
	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	preIdx := 0
	m := make(map[rune]int)
	for i, c := range s {
		if n, ok := m[c]; ok {
			if n < preIdx {
				maxLen = max(maxLen, i-preIdx+1)
			} else {
				preIdx = n + 1
				maxLen = max(maxLen, i-preIdx-1)
			}
		} else {
			maxLen = max(maxLen, i-preIdx+1)
		}
		m[c] = i
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

