## Go 实现数据结构



### 1. 栈

```go
package main

import "fmt"

type ListNode struct{ // 链表节点定义
	Val int
	Next *ListNode
}

type Stack struct {
	Node *ListNode
	Size int
}
func create() *Stack {
	node := ListNode{}
	return &Stack{&node, 0}
}

func (s *Stack) Push(val int) bool{
	tmp := s.Node
	for i:=0; i<s.Size; i++{
		tmp = tmp.Next
	}
	nd := ListNode{val, nil}
	tmp.Next = &nd
	s.Size = s.Size+1
	return true
}

func (s *Stack) Pop() int{
	if s.Size <= 0 {
		return -1
	}
	tmp := s.Node
	for i:=0; i<s.Size-1; i++{
		tmp = tmp.Next
	}
	node := tmp.Next
	tmp.Next = nil
	s.Size = s.Size-1
	return (*node).Val
}

func main() {

	stack := create()
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Push(3)
	stack.Pop()
	stack.Push(4)
	fmt.Println(stack.Pop())
}

```



