/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 *
 * https://leetcode.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (44.59%)
 * Likes:    1364
 * Dislikes: 9922
 * Total Accepted:    410.9K
 * Total Submissions: 921.4K
 * Testcase Example:  '1'
 *
 * The count-and-say sequence is the sequence of integers with the first five
 * terms as following:
 *
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 is read off as "one 1" or 11.
 * 11 is read off as "two 1s" or 21.
 * 21 is read off as "one 2, then one 1" or 1211.
 *
 * Given an integer n where 1 ≤ n ≤ 30, generate the n^th term of the
 * count-and-say sequence. You can do so recursively, in other words from the
 * previous member read off the digits, counting the number of digits in groups
 * of the same digit.
 *
 * Note: Each term of the sequence of integers will be represented as a
 * string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "1"
 * Explanation: This is the base case.
 *
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "1211"
 * Explanation: For n = 3 the term was "21" in which we have two groups "2" and
 * "1", "2" can be read as "12" which means frequency = 1 and value = 2, the
 * same way "1" is read as "11", so the answer is the concatenation of "12" and
 * "11" which is "1211".
 *
 *
 */

// @lc code=start
func countAndSay(n int) string {
	s := "1"
	if n == 1 {
		return s
	}

	for i := 2; i <= n; i++ {
		var b strings.Builder
		k := 1
		j := 0
		for j = 0; j < len(s); j++ {
			if j == 0 {
				continue
			}
			if s[j] == s[j-1] {
				k++
			} else {
				b.WriteByte(byte('0' + k))
				b.WriteByte(s[j-1])
				k = 1
			}
		}
		b.WriteByte(byte('0' + k))
		b.WriteByte(s[j-1])
		s = b.String()
	}
	return s
}

// @lc code=end

