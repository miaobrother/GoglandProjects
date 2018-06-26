package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {

	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	nodes := []treeNode{
		{value: 3},
		{},
		{7, nil, nil},
	}
	fmt.Println(nodes)
}
