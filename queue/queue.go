package main

import (
	"fmt"
)

type Node interface{}

type Queue []*Node

func (q *Queue) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n *Node) {
	if len(*q) > 0 {
		n = (*q)[0]
	} else {
		n = nil
	}
	*q = (*q)[1:]
	return n
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) String() string {
	repr := ""

	for _, c := range *q {
		repr = repr + (*c).(string) + " - "
	}
	return repr
}

func main() {
	q := Queue{}
	a := Node("Hello")
	b := Node("Bit")
	q.Push(&a)
	q.Push(&b)
	fmt.Printf("%v \n", q.String())
	q.Push(&a)
	q.Push(&a)
	fmt.Printf("%v \n", q.String())
	q.Pop()
	q.Pop()
	fmt.Printf("%v \n", q.String())
}
