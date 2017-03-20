# go-avltree
Golang implementation of an [AVL Tree](https://en.wikipedia.org/wiki/AVL_tree). An AVL tree is a [self-balancing binary search tree](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree).

Each node in the tree has a key and a value which are currently implemented as integers. It supports the following methods: Add, Remove, Update, Search, DisplayInOrder. When adding a key that exists its value is updated with the new one.

## Installation
`$ go get github.com/karask/go-avltree`

## Example usage
```
package main

import (
    "fmt"
    "github.com/karask/go-avltree"
)


func main() {
    tree := new(avltree.AVLTree)

    keys := []int{3,2,4,1,5}
    for _, key := range keys {
        tree.Add(key, key*key)
    }   

    tree.Remove(2)
    tree.Update(5, 6, 6*6)
    tree.DisplayInOrder()

    val := tree.Search(3).Value
}

```

## Notes
This code has not been thoroughly tested and is not production-ready; only basic error handling, no testing coverage, no profiling or code analysis.
