/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	hashMap := map[byte]int{}
	// cur用来记录当前最后一个重复字母的前一次出现的位置的后一位
	cur := 0
	for i := 0; i < len(s); i++ {
		if v, ok := hashMap[s[i]]; ok {
			// 这个比较是以为cur只能前进不能回头, 因为hashMap[s[i]]有可能是一个在cur之前的index的值
			if v+1 > cur {
				cur = v + 1
			}
		}

		// 计算当前的子字符串长度
		if i-cur+1 > maxLen {
			maxLen = i - cur + 1
		}

		hashMap[s[i]] = i
	}
	return maxLen
}

func main() {
	max := lengthOfLongestSubstring("abeeeecabcbb")
	fmt.Println(max)
}
