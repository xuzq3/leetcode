/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.73%)
 * Likes:    1845
 * Dislikes: 129
 * Total Accepted:    448.4K
 * Total Submissions: 979.8K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 * 
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 * 
 * 
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) < 3 {
		return 0
	}
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
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
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			} else {
				if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
					closest = sum
				}
				if sum < target {
					moveLeft()
				} else {
					moveRight()
				}
			}
		}
	}
	return closest
}
// @lc code=end

