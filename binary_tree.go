package algorithm

import "fmt"

type biTreeNode struct{
	data int
	left *biTreeNode
	right *biTreeNode
}

type biTree *biTreeNode

func CreateBiTree() (biTree, error){
	var input int
	fmt.Println("CreateBiTree, plz input:")
	num, err := fmt.Scanf("%d", &input)
	if err !=nil{
		return nil, err
	}
	fmt.Println("input:", input, "num:", num)
	if input == 255 {
		return nil, nil
	}
	var t biTreeNode
	t.data = input
	t.left, _ = CreateBiTree()
	t.right, _ = CreateBiTree()
	return &t, nil
}

func PreBiTree(t biTree){
	if t != nil {
		fmt.Println("t.data=", t.data)
		PreBiTree(t.left)
		PreBiTree(t.right)
	}
}
