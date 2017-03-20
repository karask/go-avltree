package avltree

import (
    "fmt"
)

type AVLNode struct {
    key int
    value int
    // height counts nodes (not edges)
    height int
    left *AVLNode
    right *AVLNode
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func height(node *AVLNode) int {
    if node == nil {
        return 0
    }
    return node.height
}

func recalculateHeight(node *AVLNode) {
    node.height = 1 + max(height(node.left), height(node.right))
}

func add(node *AVLNode, key int, value int) *AVLNode {
    if node == nil {
        return &AVLNode{key, value, 1, nil, nil}
    }

    if key < node.key {
        node.left = add(node.left, key, value)
    } else if key > node.key {
        node.right = add(node.right, key, value)
    } else {
        // if same key exists update value
        node.value = value
    }
    return rebalanceTree(node)
}


func rebalanceTree(node *AVLNode) *AVLNode {
    if node == nil {
        return node
    }
    recalculateHeight(node)

    // check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
    balanceFactor := height(node.left) - height(node.right)
    if balanceFactor == -2 {
        // check if child is left-heavy and rotateRight first
        if height(node.right.left) > height(node.right.right) {
            node.right = rotateRight(node.right)
        }
        return rotateLeft(node)
    } else if balanceFactor == 2 {
        // check if child is right-heavy and rotateLeft first
        if height(node.left.right) > height(node.left.left) {
            node.left = rotateLeft(node.left)
        }
        return rotateRight(node)
    }
    return node
}

func rotateLeft(node *AVLNode) *AVLNode {
    newRoot := node.right
    node.right = newRoot.left
    newRoot.left = node

    recalculateHeight(node)
    recalculateHeight(newRoot)
    return newRoot
}

func rotateRight(node *AVLNode) *AVLNode {
    newRoot := node.left
    node.left = newRoot.right
    newRoot.right = node

    recalculateHeight(node)
    recalculateHeight(newRoot)
    return newRoot
}


func search(node *AVLNode, key int) *AVLNode {
    if node == nil {
        return nil
    }
    if key < node.key {
        return search(node.left, key)
    } else if key > node.key {
        return search(node.right, key)
    } else {
        return node
    }
}


func remove(node *AVLNode, key int) *AVLNode {
    if node == nil {
        return nil
    }
    if key < node.key {
        node.left = remove(node.left, key)
    } else if key > node.key {
        node.right = remove(node.right, key)
    } else {
        if node.left != nil && node.right != nil {
           // node to delete found with both children;
           // replace values with smallest node of the right sub-tree
           rightMinNode := findSmallest(node.right)
           node.key = rightMinNode.key
           node.value = rightMinNode.value
           // delete smallest node that we replaced
           //node.right = remove(node.right, rightMinNode.key)
           node.right = remove(node.right, rightMinNode.key)
        } else if node.left != nil {
           // node only has left child
           node = node.left
        } else if node.right != nil {
           // node only has right child
           node = node.right
        } else {
           // node has no children
           node = nil
           return node
        }

    }
    return rebalanceTree(node)
}


func findSmallest(node *AVLNode) *AVLNode {
    if node.left != nil {
        return findSmallest(node.left)
    } else {
        return node
    }
}


func displayTreeInOrder(node *AVLNode) {
    if node.left != nil {
        displayTreeInOrder(node.left)
    }
    fmt.Print(node.key, " ")
    if node.right != nil {
        displayTreeInOrder(node.right)
    }
}



