package aoc2023

import (
	"reflect"
	"sync"
	"testing"
)

func TestRecursiveCount(t *testing.T) {
	tests := []struct {
		input       []int
		index       int
		expectedMap map[int]int
	}{
		{[]int{2, 1, 0, 2, 1, 0}, 0, map[int]int{0: 1, 1: 1, 2: 1, 3: 1, 4: 1}},
		{[]int{2, 2, 0, 1, 0, 1, 0}, 0, map[int]int{0: 1, 1: 1, 2: 1, 3: 1, 4: 1}},
		{[]int{1, 1, 1, 0}, 0, map[int]int{0: 1, 1: 1, 2: 1, 3: 1}},
	}

	for _, tt := range tests {
		numCopies := make(map[int]int)
		var mu sync.Mutex
		var wg sync.WaitGroup

		wg.Add(1)
		go recursiveCount(tt.input, tt.index, &numCopies, &mu, &wg)
		wg.Wait()

		mu.Lock()
		if !reflect.DeepEqual(numCopies, tt.expectedMap) {
			t.Errorf("recursiveCount(%v, %d) = %v, want %v", tt.input, tt.index, numCopies, tt.expectedMap)
		}
		mu.Unlock()
	}
}
