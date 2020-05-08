/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (45.37%)
 * Likes:    3523
 * Dislikes: 388
 * Total Accepted:    569.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 * 
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * Note:
 * 
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 * 
 */

// @lc code=start
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}

	m := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	dfs(digits, m, &res, "")
	return res
}

func dfs(digits string, m []string, res *([]string), cur string) {
	if digits == "" {
		*res = append(*res, cur)
		return
	}
	bs := m[digits[0]-'0']
	for i := 0; i < len(bs); i++ {
		dfs(digits[1:], m, res, cur+string(bs[i]))
	}
}
// @lc code=end

