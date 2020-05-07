/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.91%)
 * Likes:    6172
 * Dislikes: 745
 * Total Accepted:    840.5K
 * Total Submissions: 3.2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate triplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * 
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// to prevent the repeat
			continue
		}
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1

		moveLeft := func() {
			left++
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		}
		moveRight := func() {
			right--
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				moveLeft()
				moveRight()
			} else if sum > target {
				moveRight()
			} else {
				moveLeft()
			}
		}
	}
	return res
}
// @lc code=end

