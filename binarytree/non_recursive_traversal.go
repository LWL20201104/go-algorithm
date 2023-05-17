package binarytree

import (
	"fmt"
	"log"
)

// 中 -> 左 -> 右
func PreNonRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	log.Printf("Pre-non-recursive-traversal: ")
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		log.Printf("%d ", node.Value)
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
	}
}

// 左 -> 中 -> 右
func InNonRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	log.Printf("In-non-recursive-traversal: ")
	nodes := []*TreeNode{}
	for len(nodes) > 0 || root != nil {
		if root != nil {
			nodes = append(nodes, root)
			root = root.Left
		} else {
			root = nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			log.Printf("%d ", root.Value)
			root = root.Right
		}
	}
}

// 左 -> 右 -> 中
func PostNonRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	log.Printf("Post-non-recursive-traversal: ")
	tmpNodes := []*TreeNode{root}
	ansNodes := []*TreeNode{}
	for len(tmpNodes) != 0 {
		node := tmpNodes[len(tmpNodes)-1]
		tmpNodes = tmpNodes[:len(tmpNodes)-1]
		ansNodes = append(ansNodes, node)
		if node.Left != nil {
			tmpNodes = append(tmpNodes, node.Left)
		}
		if node.Right != nil {
			tmpNodes = append(tmpNodes, node.Right)
		}
	}
	for idx := len(ansNodes) - 1; idx >= 0; idx-- {
		log.Printf("%d ", ansNodes[idx].Value)
	}
}

// 二叉树某个节点中序遍历的后继节点
func GetFollowUpTreeNode(node *TreeNodeWithParent) {
	if node == nil {
		return
	}
	fmt.Printf("Follow-up-node: ")
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		fmt.Printf("%d \n", node.Value)
		return
	}
	for node.Parent != nil {
		if node.Parent.Left == node {
			fmt.Printf("%d \n", node.Parent.Value)
			return
		}
		node = node.Parent
	}
	fmt.Println("null")
}
