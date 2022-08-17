package leetcode_go

import "testing"

func deepestLeavesSum(root *TreeNode) int {
	ant := make([]int, 0)
	middle(root, 0, &ant)
	return ant[len(ant)-1]
}

func middle(root *TreeNode, deep int, ant *[]int) {
	if root != nil {
		if len(*ant) > deep {
			(*ant)[deep] = (*ant)[deep] + root.Val
		} else {
			*ant = append(*ant, root.Val)
		}
		middle(root.Left, deep+1, ant)
		middle(root.Right, deep+1, ant)
	}

}

func Test1302(t *testing.T) {

	root := TreeNode{Val: 1, Left: nil, Right: nil}
	left21 := TreeNode{Val: 2, Left: nil, Right: nil}
	left22 := TreeNode{Val: 3, Left: nil, Right: nil}
	left31 := TreeNode{Val: 4, Left: nil, Right: nil}
	left32 := TreeNode{Val: 6, Left: nil, Right: nil}
	left21.Left = &left31
	left22.Right = &left32
	root.Left = &left21
	root.Right = &left22
	i := deepestLeavesSum(&root)
	t.Log(i)
}
