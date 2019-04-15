/*
 * @lc app=leetcode id=9 lang=cpp
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (42.66%)
 * Total Accepted:    547.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 * 
 * Example 1:
 * 
 * 
 * Input: 121
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 * 
 * 
 * Follow up:
 * 
 * Coud you solve it without converting the integer to a string?
 * 
 */
#include <limits.h>

class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0)
            return false;

        int reverse = 0;
        int tmp = x;
        int y = 0;
        while (tmp != 0) {
            y = tmp % 10;
            tmp = tmp / 10;

            if (reverse > INT_MAX / 10)
                return false;
            else if (reverse == INT_MAX / 10 && y > INT_MAX - INT_MAX / 10 * 10)
                return false;
            
            reverse = reverse * 10 + y;
        }

        if (reverse == x)
            return true;
        else
            return false;
    }
};

