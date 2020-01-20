package algorithm

import "testing"

func TestQuick(t *testing.T){
	t.Log("test quick sort")
	quickSort(s1, 0, len(s1)-1)
	t.Logf("after quick sort, s1:%v\n", s1)
}
