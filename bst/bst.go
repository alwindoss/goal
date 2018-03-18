package bst

import (
	"log"
)

// Node holds the node of the Binary Search Tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewBST is a factory to create a BinarySearchTree
func NewBST() *BinarySearchTree {
	bst := &BinarySearchTree{Root: nil}
	return bst
}

// BinarySearchTree is the struct that holds
type BinarySearchTree struct {
	Root *Node
}

// Insert inserts the value into BinarySearchTree
func (bst *BinarySearchTree) Insert(val int) {
	node := &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
	if bst.Root == nil {
		log.Printf("root is nil")
		bst.Root = node
		return
	}
	log.Printf("I AM HERE: %d", bst.Root.Value)
	insertNode(bst.Root, node)
}

func insertNode(node *Node, newNode *Node) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}
