/*
 * @lc app=leetcode id=20 lang=cpp
 *
 * [20] Valid Parentheses
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
