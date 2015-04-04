package sort

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

func Heapsort(list []int) {
	heapify(list)
	fmt.Println(list)
	for end := len(list) - 1; end >= 0; {
		list[end], list[0] = list[0], list[end]
		end--
		siftDown(list, 0, end)
	}
}
