package main

import (
	"fmt"
	"../../algorithm"
)

func main() {
	fmt.Println("main...")
	ret, err := algorithm.CreateBiTree()
	if err != nil {
		fmt.Println(err)
	}
	algorithm.PreBiTree(ret)
}
