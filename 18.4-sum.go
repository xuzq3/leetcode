/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (32.81%)
 * Likes:    1690
 * Dislikes: 306
 * Total Accepted:    314K
 * Total Submissions: 952.6K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate quadruplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 * 
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 * 
 * 
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// to prevent the repeat
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				// to prevent the repeat
				continue
			}
			tmp := target - nums[i] - nums[j]
			left := j + 1
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
				if sum == tmp {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					moveLeft()
					moveRight()
				} else if sum > tmp {
					moveRight()
				} else {
					moveLeft()
				}
			}
		}

	}
	return res
}
// @lc code=end

