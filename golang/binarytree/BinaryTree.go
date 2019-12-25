package binarytree

import "fmt"

// 二叉树
// 定义一个节点的数据结构
type Node struct {
	data   int
	parent *Node // 父节点
	left   *Node // 左节点
	right  *Node // 右节点
}

// 建立一个BinaryTree
type BinaryTree struct {
	root *Node // 遍历节点
}

// 插入一个item
func (tree *BinaryTree) InsertItem(i int) {

	// 当根节点为空的时候 说明这个这个二叉树是没有数据的
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}

	// 2.表示是有子节点的情况
	currentNode := tree.root
	for {
		// 右节点数据
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			// 左节点数据
			if currentNode.left == nil {
				currentNode.left = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}

// 查找一个节点
func (tree *BinaryTree) SearchItem(i int) (*Node, bool) {

	// 实现逻辑
	if tree.root == nil {
		// 说明就是一个空节点 直接返回 nil false
		return nil, false
	}
	currentNode := tree.root
	// 当前节点有值
	for currentNode != nil {
		// 如果正好 待查找的值 在当前的节点 这样我们就可以 返回了
		if i == currentNode.data {
			return currentNode, true
		} else if i > currentNode.data {
			// 遍历右节点 直到这个值等于当前节点 currentNode
			currentNode = currentNode.right
		} else if i < currentNode.data {
			// 遍历左节点 直到这个值等于当前节点 currentNode
			currentNode = currentNode.left
		}
	}
	return nil, false
}

// 打印节点
func (tree *BinaryTree) PrintItem(subtree *Node) {
	if subtree.left != nil {
		tree.PrintItem(subtree.left)
	}
	// 打印
	fmt.Println(subtree.data)
	if subtree.right != nil {
		tree.PrintItem(subtree.right)
	}
}

// 中间 有序遍历
func (tree *BinaryTree) InorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.InorderTraversal(subtree.left, callback)
	}
	// 打印
	callback(subtree.data)
	if subtree.right != nil {
		tree.InorderTraversal(subtree.right, callback)
	}
}

// 前序遍历
func (tree *BinaryTree) PreorderTraversal(subtree *Node, callback func(int)) {
	callback(subtree.data)
	if subtree.left != nil {
		tree.PreorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PreorderTraversal(subtree.right, callback)
	}
}

// 后续遍历
func (tree *BinaryTree) PostorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.PostorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PostorderTraversal(subtree.right, callback)
	}
	callback(subtree.data)
}
