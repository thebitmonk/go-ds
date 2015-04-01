package main

import (
	"fmt"
)

type TrieNode struct {
	value    rune
	children []*TrieNode
	terminal bool
}

type Trie struct {
	root  *TrieNode
	count int
}

func (t *Trie) newNode(val rune) *TrieNode {
	node := new(TrieNode)
	node.value = val
	node.terminal = false
	node.children = make([]*TrieNode, 0, 0)
	return node
}

func (t *Trie) search(s string) bool {
	fmt.Printf("Searching %q \n", s)
	current := t.root
	exists := false
	for _, val := range s {
		exists = false
		for _, child := range current.children {
			if child.value == val {
				current = child
				exists = true
				break
			}
		}
		if exists == false {
			return false
		}
	}

	return current.terminal

}

func (t *Trie) existsOrAdd(s string) {
	current := t.root
	exists := false
	for _, val := range s {

		fmt.Println(val)
		for _, child := range current.children {
			if child.value == val {
				current = child
				exists = true
				break
			}
		}

		if exists {
			exists = false
			continue
		} else {
			node := t.newNode(val)
			current.children = append(current.children, node)
			current = node
			t.count = t.count + 1
		}

	}
	current.terminal = true
}

func NewTrie() *Trie {
	return &Trie{new(TrieNode), 0}
}

func main() {
	t := NewTrie()
	t.existsOrAdd("Bit")
	t.existsOrAdd("Bitmonk")
	fmt.Println(t.search("Bit"))
	fmt.Println(t.search("Bi"))
	fmt.Println(t.count)
}
