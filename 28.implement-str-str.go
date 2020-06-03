/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (33.98%)
 * Likes:    1489
 * Dislikes: 1838
 * Total Accepted:    643.7K
 * Total Submissions: 1.9M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 * 
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 * 
 * Example 1:
 * 
 * 
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 * 
 * 
 * Clarification:
 * 
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 * 
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 * 
 */

// @lc code=start
func getNext(needle string) []int {
	next := make([]int, len(needle))
	next[0] = -1

	j := 0
	k := -1
	for j < len(needle)-1 {
		if k == -1 || needle[j] == needle[k] {
			j++
			k++
			if needle[j] == needle[k] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}
	}

	return next
}

/**
 * KMP 算法
 */
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	next := getNext(needle)
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
			if j == -1 {
				i++
				j++
			}
		}
	}
	if j == len(needle) {
		return i - j
	} else {
		return -1
	}
}
// @lc code=end

