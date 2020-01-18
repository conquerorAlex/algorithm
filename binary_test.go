package main

import (
	"fmt"
	"testing"
)

var arrary = []int{1, 2, 3, 4, 5, 6}
func TestBinary(t *testing.T){
	t.Log("testBinary")
	t.Log(fmt.Sprintf("arrary:%v", arrary))
	ret, err := binary(arrary, 1)
	if err != nil{
		t.Fatal(err)
	}
	t.Logf("found successfully, ret=%d, value=%d", ret, arrary[ret])

}

func TestBinaryErr(t *testing.T){
	t.Log("testBinary")
	t.Log(fmt.Sprintf("arrary:%v", arrary))
	ret, err := binary(arrary, 0)
	if err != nil{
		t.Fatal(err)
	}
	t.Logf("found successfully, ret=%d, value=%d", ret, arrary[ret])
}