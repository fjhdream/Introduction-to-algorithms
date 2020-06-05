package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Girl struct {
	in  int
	out int
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// type GirlHeap []Girl

// func (girls GirlHeap) Len() int           { return len(girls) }
// func (girls GirlHeap) Less(i, j int) bool { return girls[i].out < girls[j].out }
// func (girls GirlHeap) Swap(i, j int)      { girls[i], girls[j] = girls[j], girls[i] }

// func (girls *GirlHeap) Push(h interface{}) {
// 	*girls = append(*girls, h.(Girl))
// }
// func (girls *GirlHeap) Pop() (x interface{}) {
// 	n := girls.Len()
// 	x = (*girls)[n-1]
// 	*girls = (*girls)[:n-1]
// 	return x
// }

func maxInOutGirls(arr [][]int) int {
	girls := []Girl{}
	gh := &IntHeap{}
	for _, inOut := range arr {
		girls = append(girls, Girl{inOut[0], inOut[1]})
	}

	sort.Slice(girls, func(i, j int) bool {
		return girls[i].in < girls[j].in
	})
	fmt.Println(girls)
	heap.Init(gh)
	ans := 0

	for i := 0; i < len(girls); i++ {

		heap.Push(gh, girls[i].out)
		for (*gh)[0] <= girls[i].in {
			heap.Pop(gh)
		}

		if ans < gh.Len() {
			ans = gh.Len()
		}
	}
	return ans
}

func main() {
	arr := [][]int{[]int{0, 30}, []int{15, 20}, []int{5, 18}, []int{4, 8}, []int{6, 7}, []int{16, 19}, []int{16, 18}, []int{16, 17}}
	ans := maxInOutGirls(arr)
	fmt.Println(ans)
}
