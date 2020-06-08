/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (31.96%)
 * Likes:    3164
 * Dislikes: 1138
 * Total Accepted:    357.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 * 
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 * 
 * The replacement must be in-place and use only constant extra memory.
 * 
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 * 
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 * 
 */

// @lc code=start

/* 思路：
 * 从右往左开始遍历，当nums[i]比nums[i+1]小时，需要替换nums[i]，
 * 找到nums[i+1:]中大于等于nums[i]的最小值nums[j]，
 * 替换nums[i]和nums[j]，
 * 最后nums[i+1:]反转，即可得出最后结果
 */
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := i + 1; j < len(nums); j++ {
				if j == len(nums)-1 || nums[i] >= nums[j+1] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			break
		}
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
// @lc code=end

