package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func mergeSortWithInsert(array []int, p, r, k int) {
	if r-p < k {
		insertionSort(array[p:r])
		return
	}
	q := int(math.Floor(float64((p + r) / 2)))
	if p < r && q < r && p < q {
		mergeSortWithInsert(array, p, q, k)
		mergeSortWithInsert(array, q, r, k)
		merge(array, p, q, r)

	}

}

//在某个临界值下插入排序比归并排序快,可以此设置k值,测试发现在200以下insertSort速度更快
func main() {
	array1 := rand.Perm(10000)
	array2 := append([]int{}, array1...)
	timeStart1 := time.Now()
	mergeSortWithInsert(array1, 0, len(array1), 200)
	//insertionSort(array1)
	timeRunning1 := time.Since(timeStart1)
	fmt.Printf("the time cost of  mergeSortWithInsert is: %d ns \n", timeRunning1)

	timeStart2 := time.Now()
	mergeSortWithoutGo(array2, 0, len(array2))
	timeRunning2 := time.Since(timeStart2)
	fmt.Printf("the time cost of  mergeSortWithoutGo  is: %d ns \n", timeRunning2)
}
