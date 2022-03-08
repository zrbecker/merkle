package merkle

import (
	"fmt"
)

func InsertBST(node Node, key string, value string) (Node, LeafNode) {
	switch n := node.(type) {
	case *leaf:
		newLeaf := NewLeaf(key, value)
		if key == n.key {
			return newLeaf, newLeaf.(*leaf)
		} else if key < n.key {
			return NewInnerNode(newLeaf, n), newLeaf.(*leaf)
		} else {
			return NewInnerNode(n, newLeaf), newLeaf.(*leaf)
		}
	case *innerNode:
		if key < n.right.MinKey() {
			newLeft, newLeaf := InsertBST(n.left, key, value)
			var newNode Node
			if newLeft.Height()-n.right.Height() > 1 {
				if newLeft.Height()-newLeft.(*innerNode).left.Height() == 1 {
					newNode = rotateRight(NewInnerNode(newLeft, n.right))
				} else {
					newLeft = rotateLeft(newLeft)
					newNode = rotateRight(NewInnerNode(newLeft, n.right))
				}
			} else {
				newNode = NewInnerNode(newLeft, n.right)
			}
			return newNode, newLeaf
		} else {
			newRight, newLeaf := InsertBST(n.right, key, value)
			var newNode Node
			if newRight.Height()-n.left.Height() > 1 {
				if newRight.Height()-newRight.(*innerNode).right.Height() == 1 {
					newNode = rotateLeft(NewInnerNode(n.left, newRight))
				} else {
					newRight = rotateRight(newRight)
					newNode = rotateLeft(NewInnerNode(n.left, newRight))
				}
			} else {
				newNode = NewInnerNode(n.left, newRight)
			}
			return newNode, newLeaf
		}
	default:
		newLeaf := NewLeaf(key, value)
		return newLeaf, newLeaf.(*leaf)
	}
}

func FindBST(node Node, key string) LeafNode {
	switch n := node.(type) {
	case *leaf:
		return n
	case *innerNode:
		if key < n.right.MinKey() {
			return FindBST(n.left, key)
		} else {
			return FindBST(n.right, key)
		}
	default:
		return nil
	}
}

func PrintBST(node Node) {
	switch n := node.(type) {
	case *leaf:
		fmt.Printf("%s: %s\n", n.key, n.value)
	case *innerNode:
		PrintBST(n.left)
		PrintBST(n.right)
	}
}

func CheckBalance(node Node) bool {
	switch n := node.(type) {
	case *leaf:
		return true
	case *innerNode:
		diff := n.left.Height() - n.right.Height()
		return diff == -1 || diff == 0 || diff == 1
	default:
		return false
	}
}

func rotateLeft(node Node) Node {
	n := node.(*innerNode)
	nr := n.right.(*innerNode)
	return NewInnerNode(
		NewInnerNode(n.left, nr.left),
		nr.right,
	)
}

func rotateRight(node Node) Node {
	n := node.(*innerNode)
	nl := n.left.(*innerNode)
	return NewInnerNode(
		nl.left,
		NewInnerNode(nl.right, n.right),
	)
}
