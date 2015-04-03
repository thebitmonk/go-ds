package main

import (
	"fmt"
)

func heapify(list []int) {
	size := len(list)
	for start := (size - 2) / 2; start >= 0; start-- {
		siftDown(list, start, size-1)
	}
}

func siftDown(list []int, start int, end int) {
	for root := start; (2*root)+1 <= end; {
		lchild := (2 * root) + 1
		rchild := lchild + 1

		idx := root
		if list[lchild] > list[idx] {
			idx = lchild
		}

		if rchild <= end && list[rchild] > list[idx] {
			idx = rchild
		}

		if idx != root {
			list[root], list[idx] = list[idx], list[root]
			root = idx
		} else {
			return
		}
	}
}

func heapsort(list []int) {
	heapify(list)
	fmt.Println(list)
	for end := len(list) - 1; end >= 0; {
		list[end], list[0] = list[0], list[end]
		end--
		siftDown(list, 0, end)
	}
}

func main() {
	list := make([]int, 10, 10)
	list = []int{50, 30, 20, 4, 5, 1, 9, 10, 11, 3}
	heapsort(list)
	fmt.Println(list)

	list = []int{3, 3333, 33, 33333, 333, 333333333, 33333333, 3333333, 333333, 3333333333}
	heapsort(list)
	fmt.Println(list)
}
