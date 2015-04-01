package main

import (
	"fmt"
	"sort"
)

type Node struct {
	freq  int
	data  string
	left  *Node
	right *Node
}

type Huffman struct {
	nodes []*Node
	root  *Node
}

type ByFrequency []*Node

//ByFrequency implements sort interface
func (a ByFrequency) Len() int {
	return len(a)
}

func (a ByFrequency) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByFrequency) Less(i, j int) bool {
	return a[i].freq < a[j].freq
}

func (h *Huffman) Pop() (*Node, *Node) {
	if len(h.nodes) >= 2 {
		sort.Sort(ByFrequency(h.nodes))
		lastMin := h.nodes[0]
		secondLastMin := h.nodes[1]
		h.nodes = h.nodes[2:]
		return lastMin, secondLastMin
	}
	return nil, nil
}

func (h *Huffman) Push(a *Node, b *Node) {
	if a == nil || b == nil {
		return
	}

	n := NewNode(a.freq+b.freq, a.data+b.data)
	n.left = a
	n.right = b

	h.nodes = append(h.nodes, n)
	h.root = n
	fmt.Println(n.data)
	fmt.Println(n.freq)
}

func (h *Huffman) traverse(path string, root *Node) {
	if root == nil {
		return
	}

	h.traverse(path+"0", root.left)
	if root.left == nil && root.right == nil {
		fmt.Println("Node " + root.data)
		fmt.Println("Path " + path)
		fmt.Println(root.freq)
	}
	h.traverse(path+"1", root.right)
}

func NewNode(freq int, data string) *Node {
	return &Node{freq, data, nil, nil}
}

func NewHuffman() *Huffman {
	return &Huffman{make([]*Node, 0, 0), new(Node)}
}

func main() {
	h := NewHuffman()
	ip := map[string]int{
		"a": 5,
		"b": 9,
		"c": 12,
		"d": 13,
		"e": 16,
		"f": 45,
	}

	for k, v := range ip {
		n := NewNode(v, k)
		h.nodes = append(h.nodes, n)
	}

	for i := 0; i < len(ip); i = i + 1 {
		a, b := h.Pop()
		h.Push(a, b)
	}

	h.traverse("", h.root)
}
