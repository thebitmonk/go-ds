package sort

import (
	"testing"
)

func isSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func TestBasicSort(t *testing.T) {
	tests := [][]int{
		{100, 20, 300, 45, 86, 78},
		{99, 10, 301, 45, 86, 78},
	}

	for _, list := range tests {
		Heapsort(list)
		if !isSorted(list) {
			t.Errorf("List is not sorted %v", list)
		}
	}

}
