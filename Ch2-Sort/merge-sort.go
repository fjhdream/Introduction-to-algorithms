package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func merge(array []int, p, q, r int) {
	intMax := int(^uint(0) >> 1)
	leftArr := append([]int{}, array[p:q]...)
	rightArr := append([]int{}, array[q:r]...)
	leftArr = append(leftArr, intMax)
	rightArr = append(rightArr, intMax)

	i := 0
	j := 0
	for k := p; k < r; k++ {
		if leftArr[i] <= rightArr[j] {
			array[k] = leftArr[i]
			i = i + 1
		} else {
			array[k] = rightArr[j]
			j = j + 1
		}
	}

}

func mergeSort(array []int, p, r int) {
	q := int(math.Floor(float64((p + r) / 2)))
	if p < r && q < r && p < q {
		go mergeSort(array, p, q)
		go mergeSort(array, q, r)
		go merge(array, p, q, r)
	} else {
		return
	}

}

func main() {

	array := rand.Perm(1000000)
	timeStart := time.Now()
	mergeSort(array, 0, len(array))
	timeRunnig := time.Since(timeStart)
	fmt.Printf("time cost: %d ns \n", timeRunnig)
}
