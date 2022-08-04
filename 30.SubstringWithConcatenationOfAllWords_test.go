/*
	30. 串联所有单词的子串
	https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
*/
package leetcode_go

import "testing"

func findSubstring(s string, words []string) []int {
	ans := []int{}
	m, n, k := len(s), len(words), len(words[0])
	for i := 0; i <= m-(n*k); i++ {
		word_map := map[string]int{}
		for _, word := range words {
			word_map[word]++
		}
		for j := 0; j < n; j++ {
			str := s[i+j*k : i+(j+1)*k]
			if word_map[str] != 0 {
				word_map[str]--
				if word_map[str] == 0 {
					delete(word_map, str)
				}
			} else {
				break
			}
		}
		if len(word_map) == 0 {
			ans = append(ans, i)
		}

	}
	return ans
}

func TestFindSubStr(t *testing.T) {
	words := []string{"word", "good", "best", "good"}
	ans := findSubstring2("wordgoodgoodgoodbestword", words)
	t.Log(ans)
}

func findSubstring2(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}
