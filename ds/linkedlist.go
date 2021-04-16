package main

import "fmt"

type Node struct {
  value int
  next *Node
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (l *LinkedList) push(node *Node) {
  l.tail.next = node
}

func main() {
  linkedList := LinkedList{}
  node1 := &Node{value: 1}
  linkedList.head = node1

  fmt.Println(linkedList.head.value)
}