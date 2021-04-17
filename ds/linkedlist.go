package main

import "fmt"

type Node struct {
  Value int
  Next *Node
}

type LinkedList struct {
  Head *Node
  Tail *Node
  Length int
}

func (l *LinkedList) AddFirst(node *Node) {
  if l.IsEmpty() {
    l.Head = node
    l.Tail = node
  } else {
    node.Next = l.Head
    l.Head = node
  }
  l.Length++
}

func (l *LinkedList) AddLast(node *Node) {
  if l.IsEmpty() {
    l.Head = node
    l.Tail = node
  } else {
    l.Head.Next = node
    l.Tail = node
  }
  l.Length++
}

func (l *LinkedList) RemoveLast() {
  if l.IsEmpty() {
    return
  }

  current := l.Head
  for current.Next.Next != nil {
    current = current.Next
  }

  current.Next = nil
  l.Tail = current
}

func (l *LinkedList) RemoveFirst() {
  if l.IsEmpty() {
    return
  }

  l.Head = l.Head.Next
}

func (l LinkedList) IsEmpty() bool {
  if l.Head == nil {
    return true
  }
  return false
}

func main() {
  myLinkedList := LinkedList{}
  
  // true
  fmt.Println(myLinkedList.IsEmpty())

  node1 := &Node{Value: 33}
  node2 := &Node{Value: 44}
  node3 := &Node{Value: 55}
  node4 := &Node{Value: 66}

  myLinkedList.AddLast(node1) // [33]
  myLinkedList.AddLast(node2) // [33, 44]
  myLinkedList.AddFirst(node3) // [55, 33, 44]
  myLinkedList.AddFirst(node4) // [66, 55, 33, 44]

  fmt.Println(myLinkedList.Head.Value)  // 66
  fmt.Println(myLinkedList.Tail.Value) // 44
  fmt.Println(myLinkedList.Length) // 4

  myLinkedList.RemoveLast() 
  fmt.Println(myLinkedList.Tail.Value) //33

  myLinkedList.RemoveFirst()
  fmt.Println(myLinkedList.Head.Value) // 55

}