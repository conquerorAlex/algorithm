package algorithm

import "testing"

var s2 = []int{1, 2, 3, 4, 5}
func TestSwap(t *testing.T){
	t.Log("testSwap")
	swap(s2, 1, 3)
	t.Log(s2)
}



