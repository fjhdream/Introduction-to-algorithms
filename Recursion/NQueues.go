package main

import (
	"fmt"
	"math"
	"time"
)

func getCountOfSolutions(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process1(0, record, n)
}

func process1(i int, record []int, n int) int {
	if i == n {
		return 1
	}

	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, record, n)
		}
	}
	return res
}

func getCountOfSolutionsGo(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	res := make(chan int, 1)
	go process1go(0, record, n, res)
	return <-res
}

func process1go(i int, record []int, n int, res chan int) {
	if i == n {
		res <- 1
	}

	nextAns := make(chan int, 1)
	newArr := make([]int, n)
	copy(newArr, record)
	ans := 0
	for j := 0; j < n; j++ {
		if isValid(newArr, i, j) {
			newArr[i] = j
			go process1go(i+1, newArr, n, nextAns)
			ans += <-nextAns
		}
	}
	res <- ans
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || math.Abs((float64)(record[k]-j)) == math.Abs((float64)(i-k)) {
			return false
		}
	}
	return true
}

//32皇后及以下的优化
func getCountOfSolutionsUsingBitMap(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	var limit int
	if n == 32 {
		limit = -1
	} else {
		limit = (1 << n) - 1
	}

	return processWithBitMap(limit, 0, 0, 0)

}
func processWithBitMap(limit, colLimit, leftDiaLimit, rightDiaLimit int) int {
	if colLimit == limit {
		return 1
	}

	pos := limit & (^(colLimit | leftDiaLimit | rightDiaLimit))
	mostRightOne := 0
	res := 0
	for pos != 0 {
		mostRightOne = pos & (^pos + 1)
		pos = pos - mostRightOne
		res += processWithBitMap(limit, colLimit|mostRightOne,
			(leftDiaLimit|mostRightOne)<<1,
			int(uint32(rightDiaLimit|mostRightOne)>>1))

	}
	return res
}

// 为什么用go? 反而速度很慢
func getCountOfSolutionsUsingBitMapWithGO(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	var limit int
	if n == 32 {
		limit = -1
	} else {
		limit = (1 << n) - 1
	}

	res := make(chan int, 1)

	processWithBitMapWithGO(limit, 0, 0, 0, res)
	return <-res

}
func processWithBitMapWithGO(limit, colLimit, leftDiaLimit, rightDiaLimit int, tune chan int) {
	if colLimit == limit {
		tune <- 1
	}

	pos := limit & (^(colLimit | leftDiaLimit | rightDiaLimit))
	mostRightOne := 0
	res := 0

	for pos != 0 {
		mostRightOne = pos & (^pos + 1)
		pos = pos - mostRightOne
		nextChan := make(chan int)
		go processWithBitMapWithGO(limit, colLimit|mostRightOne,
			(leftDiaLimit|mostRightOne)<<1,
			int(uint32(rightDiaLimit|mostRightOne)>>1), nextChan)
		res += <-nextChan

	}
	tune <- res
}

func main() {

	N := 13

	startTime := time.Now()
	ans := getCountOfSolutions(N)
	elapse := time.Since(startTime)
	fmt.Println(ans)
	fmt.Println("Without GO Routines Using ", elapse)

	// startTime = time.Now()
	// ans = getCountOfSolutionsGo(N)
	// elapse = time.Since(startTime)
	// fmt.Println(ans)
	// fmt.Println("Use GO Routines Using ", elapse)

	startTime = time.Now()
	ans = getCountOfSolutionsUsingBitMap(N)
	elapse = time.Since(startTime)
	fmt.Println(ans)
	fmt.Println("N Queues with BitMap Using ", elapse)

	startTime = time.Now()
	ans = getCountOfSolutionsUsingBitMapWithGO(N)
	elapse = time.Since(startTime)
	fmt.Println(ans)
	fmt.Println("Use GO Routines with BitMap Using ", elapse)

	startTime = time.Now()
	chanchain := make([]chan int, 10)
	count := int(math.Pow(float64(10), float64(N/2)))
	for i := 0; i < count; i++ {
		oneChan := make(chan int, 1)
		chanchain = append(chanchain, oneChan)
	}
	smallFunc := func(one chan int, two chan int) {
		ans := <-two
		one <- ans
	}
	passFunc := func(one chan int) {
		one <- 1
	}
	for i := 0; i < count; i++ {
		go smallFunc(chanchain[i], chanchain[i+1])
	}
	go passFunc(chanchain[count-1])
	elapse = time.Since(startTime)
	fmt.Println("chain of Chan", count, " Using ", elapse)

	startTime = time.Now()
	arr := make([]int, N)
	newArr := make([]int, N)
	for i := 0; i < count; i++ {
		copy(newArr, arr)
	}
	elapse = time.Since(startTime)
	fmt.Println("The cost of ", count, "copys, Using  time:", elapse)

}
