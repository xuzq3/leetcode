/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (24.91%)
 * Likes:    814
 * Dislikes: 1203
 * Total Accepted:    172.2K
 * Total Submissions: 686K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string, s, and a list of words, words, that are all of the
 * same length. Find all starting indices of substring(s) in s that is a
 * concatenation of each word in words exactly once and without any intervening
 * characters.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar"
 * respectively.
 * The output order does not matter, returning [9,0] is fine too.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * Output: []
 * 
 * 
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	var result []int

	strLength := len(s)
	totalWords := len(words)
	if strLength == 0 || totalWords == 0 {
		return result
	}
	wordLength := len(words[0])
	if wordLength == 0 {
		return result
	}

	wordsCountMap := make(map[string]int)
	for i := 0; i < totalWords; i++ {
		wordsCountMap[words[i]]++
	}

	for i := 0; i < strLength-wordLength*totalWords+1; i++ {
		// 重置map
		visited := map[string]int{}

		// 连续匹配map中的string
		j := i
		for j < strLength-wordLength+1 {
			t := s[j : j+wordLength]
			if _, ok := wordsCountMap[t]; ok {
				visited[t]++
				if visited[t] > wordsCountMap[t] {
					break
				}
				j += wordLength
			} else {
				break
			}
		}
		// 匹配的长度等于所有单词长度之和时
		if j-i == wordLength*totalWords {
			result = append(result, i)
		}
	}

	return result
}
// @lc code=end

