package algorithm

import "fmt"


type BiTree interface{
	SetValue(data int)
	SetLeft(v int)
	SetRight(v int)
	PreBiTree()
	PostBiTree()
	InBiTree()
	Layers() int
	GetLeft()*BiTreeNode
	BreathTraverse()
}

type BiTreeNode struct{
	data int
	left, right *BiTreeNode
}

func NewBiTree() BiTree{
	return &BiTreeNode{}
}

func (node *BiTreeNode)GetLeft()*BiTreeNode{
	if node == nil{
		fmt.Println("nil.node ignored.")
		return nil
	}
	return node.left
}

func (node *BiTreeNode)SetValue(data int){
	if node == nil{
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.data = data
}

func (node *BiTreeNode)SetLeft(v int){
	if node == nil{
		fmt.Println("setting left to nil.node ignored.")
		return
	}
	node.left = &BiTreeNode{data: v}
}

func (node *BiTreeNode)SetRight(v int){
	if node == nil{
		fmt.Println("setting right to nil.node ignored.")
		return
	}
	node.right = &BiTreeNode{data: v}
}

func (node *BiTreeNode)printNodeValue(){
	fmt.Print(node.data, " ")
}

func (node *BiTreeNode)PreBiTree(){
	if node == nil{
		return
	}

	node.printNodeValue()
	node.left.PreBiTree()
	node.right.PreBiTree()
}

func (node *BiTreeNode)PostBiTree(){
	if node == nil {
		return
	}

	node.left.PostBiTree()
	node.right.PostBiTree()
	node.printNodeValue()
}

func (node *BiTreeNode)InBiTree(){
	if node == nil{
		return
	}

	node.left.InBiTree()
	node.printNodeValue()
	node.right.InBiTree()
}

func (node *BiTreeNode)Layers() int {
	if node == nil{
		return 0
	}
	leftLayers := node.left.Layers()
	rightLayers := node.right.Layers()
	if leftLayers > rightLayers{
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

func (node *BiTreeNode)BreathTraverse(){
	if node == nil{
		return
	}
	var slice []*BiTreeNode
	fmt.Print(node.data, " ")
	slice = append(slice, node.left)
	slice = append(slice, node.right)
	for {
		if len(slice) == 0{
			break
		}
		fmt.Print(slice[0].data, " ")
		if slice[0].left != nil{
			slice = append(slice, slice[0].left)
		}
		if slice[0].right != nil{
			slice = append(slice, slice[0].right)
		}
		slice = slice[1:]
	}

}