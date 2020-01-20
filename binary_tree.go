package algorithm

import "fmt"


type BiTreeNode struct{
	data int
	Left, Right *BiTreeNode
}

func (node *BiTreeNode)SetValue(data int){
	if node == nil{
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.data = data
}

func (node *BiTreeNode)SetLeft(n *BiTreeNode){
	if node == nil{
		fmt.Println("setting Left to nil.node ignored.")
		return
	}
	node.Left = n
}

func (node *BiTreeNode)SetRight(n *BiTreeNode){
	if node == nil{
		fmt.Println("setting right to nil.node ignored.")
		return
	}
	node.Right = n
}

func (node *BiTreeNode)printNodeValue(){
	fmt.Print(node.data, " ")
}

func CreateBiTree() (*BiTreeNode, error){
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
	var t BiTreeNode
	t.data = input
	t.Left, _ = CreateBiTree()
	t.Right, _ = CreateBiTree()
	return &t, nil
}

func (node *BiTreeNode)PreBiTree(){
	if node == nil{
		//fmt.Println("tree empty now")
		return
	}

	//fmt.Println("node.data=", node.data)
	node.printNodeValue()
	node.Left.PreBiTree()
	node.Right.PreBiTree()
}

func (node *BiTreeNode)PostBiTree(){
	if node == nil {
		//fmt.Println("tree empty now")
		return
	}

	node.Left.PostBiTree()
	node.Right.PostBiTree()
	//fmt.Println("node.data=", node.data)
	node.printNodeValue()
}

func (node *BiTreeNode)InBiTree(){
	if node == nil{
		//fmt.Println("tree empty now")
		return
	}

	node.Left.InBiTree()
	//fmt.Println("node.data=", node.data)
	node.printNodeValue()
	node.Right.InBiTree()
}

func (node *BiTreeNode)Layers() int {
	if node == nil{
		return 0
	}
	leftLayers := node.Left.Layers()
	rightLayers := node.Right.Layers()
	if leftLayers > rightLayers{
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}