package main

import (
	"fmt"
	"../../algorithm"
)


func main() {
	fmt.Println("main...")

	biTree := algorithm.NewBiTree()

	biTree.SetValue(1)
	biTree.SetLeft(2)
	biTree.SetRight(3)
	l := biTree.GetLeft()
	l.SetLeft(4)
	l.SetRight(5)


	fmt.Println("\npre BiTree")
	biTree.PreBiTree()

	fmt.Println("\npost BiTree")
	biTree.PostBiTree()

	fmt.Println("\nin BiTree")
	biTree.InBiTree()

	fmt.Println("\nBiTree layers:", biTree.Layers())

	fmt.Println("\nbreathTraverse")
	biTree.BreathTraverse()

	fmt.Println("\ndepthTraverse")
	biTree.DepthTraverse()
}
