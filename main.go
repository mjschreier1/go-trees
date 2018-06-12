package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	test := &Node{
		Value: 123,
	}
	printNode(test)
}

func printNode(n *Node) {
	fmt.Print("Value: ", n.Value)
	fmt.Println()
}
