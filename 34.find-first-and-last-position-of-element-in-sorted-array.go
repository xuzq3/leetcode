/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (35.36%)
 * Likes:    3342
 * Dislikes: 144
 * Total Accepted:    493.6K
 * Total Submissions: 1.4M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * If the target is not found in the array, return [-1, -1].
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * 
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	start, end := -1, -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			start = getStart(nums, left, mid, target)
			end = getEnd(nums, mid, right, target)
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{start, end}
}

func getStart(nums []int, left int, right int, target int) int {
	for left <= right {
		if left == right {
			return right
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			break
		}
	}
	return right
}

func getEnd(nums []int, left int, right int, target int) int {
	for left <= right {
		if left == right {
			return left
		}
		mid := (left+right)/2 + 1
		if nums[mid] == target {
			left = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			break
		}
	}
	return left
}
// @lc code=end

