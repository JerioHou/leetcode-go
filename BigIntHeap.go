package leetcode_go

// 大顶堆
type BigIntHeap []int

func (h BigIntHeap) Len() int {
	return len(h)
}
func (h BigIntHeap) Less(i, j int) bool {

	return h[i] > h[j]
}

func (h BigIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BigIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 取的时候,先把堆顶元素置换到末尾，然后从末尾取
// 所以每次实际取的是堆顶元素
func (h *BigIntHeap) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
