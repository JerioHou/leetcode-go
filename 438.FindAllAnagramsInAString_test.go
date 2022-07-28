package leetcode_go

import (
	"testing"
)

func findAnagrams(s string, p string) []int {
	var ans []int
	var ls, lp = len(s), len(p)
	if ls < lp {
		return ans
	}
	var sCounrt, pCount [26]int
	for i, ch := range p {
		pCount[ch-'a']++
		sCounrt[s[i]-'a']++
	}
	if sCounrt == pCount {
		ans = append(ans, 0)
	}
	for i, ch := range s[:ls-lp] {
		sCounrt[ch-'a']--
		sCounrt[s[i+lp]-'a']++
		if sCounrt == pCount {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func TestIsEctopicWord(t *testing.T) {

	ans := findAnagrams("werwerfxafd", "er")
	t.Log(ans)

}
