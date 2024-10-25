package main

/*
Given a string s, find the length of the longest
substring
 without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

import "math"

func initLastIndex() map[uint8]int {
	lastIndex := make(map[uint8]int)
	for i := 0; i < 256; i++ {
		lastIndex[uint8(i)] = -1
	}
	return lastIndex
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	result := 0
	lastIndex := initLastIndex()
	for end := 0; end < len(s); end++ {
		start = int(math.Max(float64(start), float64(lastIndex[s[end]]+1)))
		result = int(math.Max(float64(result), float64(end-start+1)))
		lastIndex[s[end]] = end
	}
	return result
}
