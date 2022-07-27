package leetcode_go

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode = new(ListNode)
	var tail = head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			tail.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return head.Next
}
