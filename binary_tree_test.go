package algorithm

import (
	"fmt"
	"testing"
)

func TestPreBiTree(t *testing.T){
	t.Log("testPreBiTree")
	ret, err := CreateBiTree()
	if err != nil {
		t.Fatal(err)
	}
	PreBiTree(ret)
}

func TestScanf(t *testing.T){
	t.Log("testScanf")
	var a int
	fmt.Scanf("%d", &a)
}