/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (26.35%)
 * Likes:    3788
 * Dislikes: 655
 * Total Accepted:    402.8K
 * Total Submissions: 1.5M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 * 
 * 
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * 
 * 
 * The matching should cover the entire input string (not partial).
 * 
 * Note:
 * 
 * 
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore,
 * it matches "aab".
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 * 
 * 
 */

// @lc code=start
/**
 * 递归
 */
 func isMatch_recursive(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	} else if len(p) == 1 {
		return len(s) == 1 && isSingleMatch(s[0], p[0])
	} else {
		if p[1] != '*' {
			return len(s) >= 1 && isSingleMatch(s[0], p[0]) && isMatch(s[1:], p[1:])
		} else {
			return isMatch(s, p[2:]) || (len(s) >= 1 && isSingleMatch(s[0], p[0]) && isMatch(s[1:], p))
		}
	}
}

/**
 * 动态规划
 */
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	// 从后往前匹配
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			if p[j] != '*' {
				dp[i][j] = i < len(s) && dp[i+1][j+1] && isSingleMatch(s[i], p[j])
			} else {
				if j <= 0 {
					dp[i][j] = false
				} else {
					j--
					dp[i][j] = dp[i][j+2] || (i < len(s) && isSingleMatch(s[i], p[j]) && dp[i+1][j])
				}
			}
		}
	}
	return dp[0][0]
}

func isSingleMatch(s byte, p byte) bool {
	if p == '.' {
		return true
	} else {
		return s == p
	}
}
// @lc code=end

