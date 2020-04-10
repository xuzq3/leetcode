/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (28.66%)
 * Likes:    6335
 * Dislikes: 967
 * Total Accepted:    626K
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * 
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 * 
 * You may assume nums1 and nums2 cannot be both empty.
 * 
 * Example 1:
 * 
 * 
 * nums1 = [1, 3]
 * nums2 = [2]
 * 
 * The median is 2.0
 * 
 * 
 * Example 2:
 * 
 * 
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 
 * The median is (2 + 3)/2 = 2.5
 * 
 * 
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	n := n1 + n2

	// to ensure n1 <= n2
	if n1 > n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}

	left := 0
	right := n1
	i := 0
	j := 0
	for {
		i = left + (right-left)/2
		j = n/2 - i

		if i > 0 && nums1[i-1] > nums2[j] {
			right = left + (right-left)/2
		} else if i < n1 && nums2[j-1] > nums1[i] {
			left = left + (right-left+1)/2
		} else {
			break
		}
	}

	getLeft := func(i, j int) int {
		if i > 0 && j > 0 {
			if nums1[i-1] > nums2[j-1] {
				return nums1[i-1]
			} else {
				return nums2[j-1]
			}
		} else if i > 0 {
			return nums1[i-1]
		} else if j > 0 {
			return nums2[j-1]
		}
		return 0
	}

	getRight := func(i, j int) int {
		if i < n1 && j < n2 {
			if nums1[i] < nums2[j] {
				return nums1[i]
			} else {
				return nums2[j]
			}
		} else if i < n1 {
			return nums1[i]
		} else if j < n2 {
			return nums2[j]
		}
		return 0
	}

	if n%2 == 0 {
		left := getLeft(i, j)
		right := getRight(i, j)
		return float64(left+right) / 2
	} else {
		middle := getRight(i, j)
		return float64(middle)
	}
}
// @lc code=end

