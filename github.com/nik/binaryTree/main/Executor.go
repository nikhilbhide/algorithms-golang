package main

import (
	"fmt"
	"github.com/nik/binaryTree/binarytree"
	"github.com/nik/binaryTree/binarytree/bst"
	"github.com/nik/binaryTree/binarytree/node"
	"github.com/nik/binaryTree/binarytree/traversal"
)

func main() {
	Numbers := [] int{10, 5, 50, 3, 20, 24, 30, 100, 90, 70, 23, 15}

	for _, item := range Numbers {
		newNode := node.Node{
			Left:  nil,
			Right: nil,
			Num:   item,
		}

		bst.Insert(&newNode)
	}
	fmt.Println("In order traversal order is as follows :")
	binarytree.DisplayTree(traversal.InorderTraversal(*bst.RootNode))
	fmt.Println("Pre order traversal order is as follows :")
	binarytree.DisplayTree(traversal.PreorderTraversal(*bst.RootNode))
	fmt.Println("Post order traversal order is as follow :")
	binarytree.DisplayTree(traversal.PostorderTraversal(*bst.RootNode))
}