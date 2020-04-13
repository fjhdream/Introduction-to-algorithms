package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			//fmt.Println("key:", key, "arry[j]:", array[j])
			array[j], array[j+1] = array[j+1], array[j]
			j = j - 1
			//fmt.Println(array)
		}
		array[j+1] = key
	}
}

// func main() {
// 	exampleArray := []int{2, 5, 4, 3, 1}

// 	timeStart := time.Now()
// 	insertionSort(exampleArray)
// 	timeRunning := time.Since(timeStart)
// 	fmt.Printf("the length of %d array sorts cost %d ns\n", len(exampleArray), timeRunning)

// 	rand.Seed(time.Now().UnixNano())
// 	exampleArray = rand.Perm(1000)
// 	timeStart = time.Now()
// 	insertionSort(exampleArray)
// 	timeRunning = time.Since(timeStart)
// 	fmt.Printf("the length of %d array sorts cost %d ns\n", len(exampleArray), timeRunning)
// }
