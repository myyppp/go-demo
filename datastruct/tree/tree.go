package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preOrderRec(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preOrderRec(root.Left)
	preOrderRec(root.Right)
}

// 递归
func midOrderRec(root *TreeNode) {
	if root == nil {
		return
	}
	midOrderRec(root.Left)
	fmt.Print(root.Val, " ")
	midOrderRec(root.Right)

}

// 递归
func postOrderRec(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderRec(root.Left)
	postOrderRec(root.Right)
	fmt.Print(root.Val, " ")
}

// 先序遍历 一
func preOrder(root *TreeNode) {
	s := new(Stack)
	for root != nil || !s.Empty() {
		for root != nil {
			s.Push(root)
			fmt.Print(root.Val, " ")
			root = root.Left
		}
		root = s.Pop()
		root = root.Right
	}
}

// 先序遍历 二
func pre(root *TreeNode) {
	if root == nil {
		return
	}
	s := new(Stack)
	s.Push(root)
	for !s.Empty() {
		node := s.Pop()
		fmt.Print(node.Val, " ")
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

// 中序遍历
// 扫描根节点的所有左节点，并入栈
// 出栈一个节点，访问它
// 扫描该节点的右节点，入栈
// 扫描该右节点的所有左节点
func midOrder(root *TreeNode) {
	s := new(Stack)
	for root != nil || !s.Empty() {
		for root != nil { // 根节点入栈，遍历左子树
			s.Push(root)
			root = root.Left
		}
		// 根节点出栈，访问根节点，遍历右子树
		root = s.Pop()
		fmt.Print(root.Val, " ")
		root = root.Right
	}
}

// 后序遍历什么时候可以访问节点？
// 1. 当前经过节点是叶子节点。
// 2. 当前经过节点的右子节点是上一次访问的节点
//
// 前序和中序遍历的时候，经过的节点都是 左 -> 根 -> 右
// 但在后序遍历时，需要经过根节点，先访问右节点
// 访问完右节点访问根节点
// 记录上次访问的节点，判断当前节点和上次访问的节点的关系
// 如果当前访问的节点的右节点是上次访问的节点，那么需要访问当前节点
func postOrder(root *TreeNode) {
	s := new(Stack)
	var lastVisit *TreeNode // 用于记录上一次访问的节点
	for root != nil || !s.Empty() {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		cur := s.Peek() // 栈顶元素
		if cur.Right == nil || cur.Right == lastVisit {
			s.Pop()
			fmt.Print(cur.Val, " ")
			lastVisit = cur // 记录上一次访问的节点
		} else {
			root = cur.Right
		}
	}
}

// 深度优先遍历 递归
func treeDFSRec(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	treeDFSRec(root.Left)
	treeDFSRec(root.Right)
}

// 深度优先遍历 非递归
func treeDFS(root *TreeNode) {
	s := new(Stack)
	for root != nil || !s.Empty() {
		for root != nil {
			s.Push(root)
			fmt.Print(root.Val, " ")
			root = root.Left
		}
		root = s.Pop()
		root = root.Right
	}
}

func treeDFSDivide(root *TreeNode) {

}

// 层次遍历
func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	list := []*TreeNode{} // 队列
	list = append(list, root)
	for len(list) != 0 {
		node := list[0]
		list = list[1:]
		fmt.Print(node.Val, " ")
		if node.Left != nil {
			list = append(list, node.Left)
		}
		if node.Right != nil {
			list = append(list, node.Right)
		}
	}
}

// 层次遍历 递归
func levelOrderRec(root *TreeNode) {
	d := depth(root)
	for level := 0; level < d; level++ {
		printLevel(root, level)
	}
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func printLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == 0 {
		fmt.Print(root.Val, " ")
	}
	printLevel(root.Left, level-1)
	printLevel(root.Right, level-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
