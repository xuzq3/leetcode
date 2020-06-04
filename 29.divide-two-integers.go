/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.20%)
 * Likes:    1107
 * Dislikes: 5031
 * Total Accepted:    269.8K
 * Total Submissions: 1.7M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division and mod operator.
 * 
 * Return the quotient after dividing dividend by divisor.
 * 
 * The integer division should truncate toward zero, which means losing its
 * fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) =
 * -2.
 * 
 * Example 1:
 * 
 * 
 * Input: dividend = 10, divisor = 3
 * Output: 3
 * Explanation: 10/3 = truncate(3.33333..) = 3.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: dividend = 7, divisor = -3
 * Output: -2
 * Explanation: 7/-3 = truncate(-2.33333..) = -2.
 * 
 * 
 * Note:
 * 
 * 
 * Both dividend and divisor will be 32-bit signed integers.
 * The divisor will never be 0.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 2^31 − 1 when the
 * division result overflows.
 * 
 * 
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	isNegative := false
	if dividend < 0 {
		isNegative = !isNegative
		dividend = -dividend
	}
	if divisor < 0 {
		isNegative = !isNegative
		divisor = -divisor
	}

	res := 0
	k := 0
	for dividend >= divisor {
		// 二分法
		k = 0
		for dividend >= (divisor << k << 1) {
			k++
		}
		res += (1 << k)
		dividend -= (divisor << k)
	}
	if isNegative {
		res = -res
	}
	return res
}
// @lc code=end

