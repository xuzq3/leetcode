/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (33.72%)
 * Likes:    4770
 * Dislikes: 447
 * Total Accepted:    708.1K
 * Total Submissions: 2.1M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * 
 */

// @lc code=start

// 直接比较target
func search(nums []int, target int) int {
	l := len(nums)
	if l <= 0 {
		return -1
	}
	left, right := 0, l-1
	for left <= right {
		mid := (right + left) / 2
		if left == right {
			if nums[mid] != target {
				return -1
			} else {
				return mid
			}
		}

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if nums[mid] >= nums[left] || target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] || nums[mid] <= nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// 算出偏移量后再二分查找
func search(nums []int, target int) int {
	l := len(nums)
	if l <= 0 {
		return -1
	}

	// 找出正常序列到变换后序列的偏移量offset
	left, right := 0, l-1
	for nums[left] > nums[right] && left < right {
		mid := (right + left) / 2
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	offset := left

	// 根据偏移量进行二分查找
	i := sort.Search(l, func(i int) bool {
		return nums[(i+offset)%l] >= target
	})
	if i < l && nums[(i+offset)%l] == target {
		return (i + offset) % l
	}
	return -1
}
// @lc code=end

