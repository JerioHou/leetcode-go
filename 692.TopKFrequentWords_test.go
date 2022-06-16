/**
 * 给一非空的单词列表，返回前 k 个出现次数最多的单词。
 *
 * 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
 *
 * 示例 1：
 *
 * 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * 输出: ["i", "love"]
 * 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
 *     注意，按字母顺序 "i" 在 "love" 之前。
 *
 *
 * 示例 2：
 *
 * 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
 * 输出: ["the", "is", "sunny", "day"]
 * 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
 *     出现次数依次为 4, 3, 2 和 1 次。
 *
 *
 * 注意：
 *
 * 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
 * 输入的单词均由小写字母组成。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/top-k-frequent-words
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package leetcode_go

import (
	"container/heap"
	"testing"
)

type pair struct {
	word  string
	count int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	// 小顶堆
	return a.count < b.count || a.count == b.count && a.word > b.word
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

// 取的时候,先把堆顶元素置换到末尾，然后从末尾取
// 所以每次实际取的是堆顶元素
// 小顶堆，取的是最小值
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func topKFrequent(words []string, k int) []string {
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	h := &hp{}
	for word, cnt := range wordMap {
		heap.Push(h, pair{word, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).word
	}
	return ans
}

func TestTopKFrequent(t *testing.T) {
	words := []string{"i", "love", "i", "coding", "love", "coding"}
	t.Log(topKFrequent(words, 2))
}
