/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.74%)
 * Likes:    4412
 * Dislikes: 203
 * Total Accepted:    905.6K
 * Total Submissions: 2.4M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */

// @lc code=start
#include <stack>
#include <string>
using namespace std;

class Solution
{
public:
    bool isValid(string s)
    {
        stack<char> sc;
        int len = s.length();
        char c;
        char t;
        for (int i = 0; i < len; i++)
        {
            c = s.at(i);
            if (c == '(' || c == '{' || c == '[')
            {
                sc.push(c);
                continue;
            }

            if (sc.empty())
            {
                return false;
            }

            t = sc.top();
            if ((c == ')' && t == '(') || (c == ']' && t == '[') || (c == '}' && t == '{'))
            {
                sc.pop();
                continue;
            }

            return false;
        }
        return sc.empty();
    }
};
// @lc code=end
