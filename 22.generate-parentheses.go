/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (60.59%)
 * Likes:    4683
 * Dislikes: 249
 * Total Accepted:    516K
 * Total Submissions: 844.1K
 * Testcase Example:  '3'
 *
 * 
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 * 
 * 
 * 
 * For example, given n = 3, a solution set is:
 * 
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 * 
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var results []string
	i, j := 0, 0
	s := ""
	generateSub(n, i, j, s, &results)
	return results
}

func generateSub(n int, i int, j int, s string, results *[]string) {
	if i == n && j == 0 {
		*results = append(*results, s)
		return
	}
	if i < n {
		s2 := s + "("
		i2 := i + 1
		j2 := j + 1
		generateSub(n, i2, j2, s2, results)
	}
	if j > 0 {
		s2 := s + ")"
		i2 := i
		j2 := j - 1
		generateSub(n, i2, j2, s2, results)
	}
}
// @lc code=end

