package bst

import (
	"github.com/nik/binaryTree/binarytree/node"
)


var RootNode *node.Node

func Insert(node *node.Node) *node.Node {
	if (RootNode == nil) {
		RootNode = node
	} else {
		currentNode := RootNode
		for ; currentNode != nil; {
			if (node.Num <= currentNode.Num) {
				if (currentNode.Left == nil) {
					currentNode.Left = node
					break
				} else {
					currentNode = currentNode.Left
				}
			} else {
				if (currentNode.Right == nil) {
					currentNode.Right = node
					break
				} else {
					currentNode = currentNode.Right
				}
			}
		}
		currentNode = node
	}
	return RootNode
}
