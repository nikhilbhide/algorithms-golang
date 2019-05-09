package binarytree

import (
	"fmt"
)
//Displays the tree node wise and follows the traversal order
func DisplayTree(traversal [] string) {
	for _, value := range traversal {
		fmt.Println("Item :", value)
	}
}
