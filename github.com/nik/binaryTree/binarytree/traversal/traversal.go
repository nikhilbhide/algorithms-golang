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
	stack := list.New()
	//mark current node to point to the root node
	var currentNode *node.Node
	currentNode =  &RootNode

	//loop until stack is empty or current node is not nil
	for ; (currentNode!=nil || stack.Len() != 0); {
		if (currentNode!=nil) {
			//push the current node on the stack
			stack.PushFront(currentNode)
			//make the current node to point to left child
			currentNode = currentNode.Left
		} else {
			//just peek at the top of the stack
			element := stack.Front()
			//remove the element from stack
			stack.Remove(element)
			//cast *Element to node.Node50
			currentNode = element.Value.(*node.Node)
			//add the current element to traversal order and remove the element from list
			traversalOrder = append(traversalOrder, strconv.Itoa(currentNode.Num))
			//make current node to point to right of the current node
			currentNode= currentNode.Right
		}
	}

	return traversalOrder
}

/*
Traverses the binary search tree in preorder traversal.
Preorder traversal means visit current node then Left node and then Right node
 */
func PreorderTraversal(RootNode node.Node) [] string {
	//variable to hold the traversed nodes
	var traversalOrder [] string

	//create a new list that is to be used as stack - last in first out
	stack := list.New()
	//push the root node to the stack
	stack.PushFront(RootNode)

	//loop until stack is empty
	for ; stack.Len() != 0; {
		//flag is used to understand whether to look for Right subnode
		var currentNode node.Node
		//just peek at the top of the stack
		element := stack.Front()
		stack.Remove(element)
		//cast *Element to node.Node
		currentNode = element.Value.(node.Node)
		//add the current element to traversal order and remove the element from list
		traversalOrder = append(traversalOrder, strconv.Itoa(currentNode.Num))

		if (currentNode.Right != nil) {
			//push the Right element on stack and set the current node to Right node of the current node
			stack.PushFront(*currentNode.Right)
		}

		if (currentNode.Left != nil) {
			//push the Left element on stack and set the current node to Left node of the current node
			stack.PushFront(*currentNode.Left)
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
	//stack to hold the visited nodes
	traversalOrderStack:= list.New()
	//create a new stack
	stack := list.New()
	//push the root node to the stack
	stack.PushFront(RootNode)

	//loop until stack is empty
	for ; stack.Len() != 0; {
		//just peek at the top of the list
		element := stack.Front()
		//cast *Element to node.Node
		currentNode:= element.Value.(node.Node)
		traversalOrderStack.PushFront(strconv.Itoa(currentNode.Num))
		stack.Remove(element)
		if (currentNode.Left != nil) {
			//push the Right element on stack
			stack.PushFront(*currentNode.Left)
		}

		if (currentNode.Right != nil) {
			//add the Left element on stack
			stack.PushFront(*currentNode.Right)
		}
	}

	for ;traversalOrderStack.Len()!=0; {
		element := traversalOrderStack.Front()
		traversalOrder = append(traversalOrder,element.Value.(string))
		traversalOrderStack.Remove(element)
	}

	return traversalOrder
}
