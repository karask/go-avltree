package main

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


func displayTreeInOrder(node *AVLNode) {
    if node.left != nil {
        displayTreeInOrder(node.left)
    }
    fmt.Print(node.key, " ")
    if node.right != nil {
        displayTreeInOrder(node.right)
    }
}


func main() {
    var root *AVLNode
    //keys := []int{3,2,4,1,5}
    //keys := []int{1,2,3,4,5,6,7,8,9}
    keys := []int{15,3,9,44,4,8,13}
    for _, key := range keys {
        root = add(root, key, key*key)
    }

    displayTreeInOrder(root)
}

