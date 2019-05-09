package traversal

import (
	"container/list"
	"github.com/nik/binaryTree/binarytree/node"
	"strconv"
)

//Traverses the binary tree in inorder traversal
//It returns the slice of string items, a slice in which each item denotes the value of node
//Inorder traversal gives the items in sorted order
func InorderTraversal(RootNode node.Node) [] string {
	//variable to hold the traversed nodes
	var traversalOrder [] string

	//create a new list that is to be used as stack - last in first out
	queue := list.New()
	//push the root node to the stack
	queue.PushFront(RootNode)

	//create a set of visited nodes. Initially it will be empty
	visited := make(map[string]bool)

	//loop until stack is empty
	for ; queue.Len() != 0; {
		//flag is used to understand whether to look for Right subnode
		flag := 0
		var currentNode node.Node

		//just peek at the top of the stack
		element := queue.Front()

		//cast *Element to node.Node
		currentNode = element.Value.(node.Node)

		if (currentNode.Left != nil && !visited[strconv.Itoa((*currentNode.Left).Num)]) {
			//push the Left element on stack and set the current node to Left node of the current node
			queue.PushFront(*currentNode.Left)
			currentNode = *currentNode.Left
		} else {
			//either the Left node is null or its already visited
			flag = 1
			//add the current node to visited list
			visited[strconv.Itoa(currentNode.Num)]=true
			//add the current element to traversal order and remove the element from stack
			traversalOrder = append(traversalOrder, strconv.Itoa(currentNode.Num))
			queue.Remove(element)
		}

		if (currentNode.Right != nil && flag==1) {
			//push the Right element on stack and set the current node to Right node of the current node
			queue.PushFront(*currentNode.Right)
			currentNode = *currentNode.Right
		}
	}

	return traversalOrder
}

/*
Traverses the binary search tree in preorder traversal.
Preorder traversal means visit current node then Left node and then Right node
 */
func PreorderTraversal(RootNode node.Node) [] string {
	var traversalOrder [] string

	//create a new list that is to be used as stack - first in first out
	//you add element to the rear side of queue and remove the element from front side
	queue := list.New()
	queue.PushBack(RootNode)

	//loop until queue is empty
	for ; queue.Len() != 0; {
		var currentNode node.Node

		//get the element from the queue
		element := queue.Front()
		currentNode = element.Value.(node.Node)
		//remove the element from queue
		queue.Remove(element)
		//node is visited and add element to the traversal list
		traversalOrder = append(traversalOrder, strconv.Itoa(currentNode.Num))

		//visit Left node
		if (currentNode.Left != nil) {
			queue.PushBack(*currentNode.Left)
		}
		//visit Right node
		if (currentNode.Right != nil) {
			queue.PushBack(*currentNode.Right)
		}
	}

	return traversalOrder
}

/*
Traverses the binary search tree in postorder traversal.
Postorder traversal means visit Left subtree, Right subtree and finally the root node of the Left and Right subtree
*/
func PostorderTraversal(RootNode node.Node) [] string {
	//variable to hold the traversed nodes
	var traversalOrder [] string

	//create a new list that is to be used as stack - last in first out
	queue := list.New()
	//push the root node to the stack
	queue.PushFront(RootNode)

	//create a set of visited nodes. Initially it will be empty
	visited := make(map[string]bool)

	//loop until stack is empty
	for ; queue.Len() != 0; {
		//flag is used to understand whether to look for root node of the visited left and right subtrees
		flag:=0

		var currentNode node.Node
		//just peek at the top of the stack
		element := queue.Front()
		//cast *Element to node.Node
		currentNode = element.Value.(node.Node)

		if (currentNode.Left != nil && !visited[strconv.Itoa((currentNode.Left).Num)]) {
			//push the Left element on stack and set the current node to Left node of the current node
			queue.PushFront(*currentNode.Left)
			currentNode = *currentNode.Left
		} else if (currentNode.Right != nil && !visited[strconv.Itoa((currentNode.Right).Num)]) {
			//push the Right element on stack and set the current node to Right node of the current node
			queue.PushFront(*currentNode.Right)
			currentNode = *currentNode.Right
		} else {
			flag = 1
		}

		if(flag==1) {
			//Left subtree and Right subtree of current node are visited
			//add the current node to visited list
			visited[strconv.Itoa(currentNode.Num)]=true
			//add the current element to traversal order and remove the element from stack
			traversalOrder = append(traversalOrder, strconv.Itoa(currentNode.Num))
			queue.Remove(element)
		}
	}

	return traversalOrder
}