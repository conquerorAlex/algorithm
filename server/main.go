package main

import (
	"fmt"
	"../../algorithm"
)

func main() {
	fmt.Println("main...")

	BiTree := &algorithm.BiTreeNode{}
	BiTree.SetValue(1)
	BiTree.SetLeft(&algorithm.BiTreeNode{})
	BiTree.SetRight(&algorithm.BiTreeNode{})
	BiTree.Left.SetValue(2)
	BiTree.Right.SetValue(3)

	fmt.Println("\npre BiTree")
	BiTree.PreBiTree()

	fmt.Println("\npost BiTree")
	BiTree.PostBiTree()

	fmt.Println("\nin BiTree")
	BiTree.InBiTree()
}
