package leetcode_go

import (
	"strings"
	"testing"
)

/*
  https://leetcode-cn.com/problems/zigzag-conversion/
*/

func TestZ(t *testing.T) {
	t.Log(zconvert("ABCDE", 4))
	//t.Log("AVCE"[1:2])
}
func zconvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	cycle := 2 * (numRows - 1)
	builder := strings.Builder{}
	// z字形的次数
	counts := n/cycle + 1
	for i := 0; i < numRows; i++ {
		for j := 0; j < counts; j++ {
			if j*cycle+i < n {
				builder.WriteString(s[j*cycle+i : j*cycle+i+1])
			}
			if i > 0 && i < numRows-1 {
				after := j*cycle + (cycle - i)
				if after < n {
					builder.WriteString(s[after : after+1])
				}
			}
		}
	}
	return builder.String()
}
