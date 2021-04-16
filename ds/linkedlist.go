package main

import "fmt"

type Node struct {
  Value int
  Next *Node
}

type LinkedList struct {
  Head *Node
  Tail *Node
}

func (l *LinkedList) push(node *Node) {
  l.Tail.Next = node
}

func main() {
  linkedList := LinkedList{}
  node1 := &Node{Value: 1}
  linkedList.Head = node1

  fmt.Println(linkedList.Head.Value)
}