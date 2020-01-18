package main

import "testing"

var s1 = []int{3, 8, 1, 5, 2, 9, 1, 3}
func TestBubble(t *testing.T){
	t.Log("testBubble")
	t.Logf("s1:%v", s1)
	t.Logf("S1[0]:%d", s1[0])
	bubbleSort(s1)
	t.Logf("after sort, s1:%v ", s1)
}
