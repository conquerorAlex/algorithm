package algorithm


/*
二叉树：每个节点最多有两个子节点，且这两个节点有左右次序之分
满二叉树： 二叉树所有非叶子节点都被填满；深度为k，节点总数为：2的k次方-1
完全二叉树： 有n个节点，深度为k二叉树，当且仅当，节点的排列序号，与深度为k的满二叉树中序号为1至n的节点一一对应，
           是一种特殊的满二叉树
堆排序：分大顶堆 和 小顶堆
*/
func swap(s []int, i, j int){
	s[i], s[j] = s[j], s[i]
}