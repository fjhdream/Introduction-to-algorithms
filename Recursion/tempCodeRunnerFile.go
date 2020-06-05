
	// startTime = time.Now()
	// chanchain := make([]chan int, 10)
	// count := 10
	// for i := 0; i < count; i++ {
	// 	oneChan := make(chan int, 1)
	// 	chanchain = append(chanchain, oneChan)
	// }
	// for i := 0; i < count; i++ {
	// 	chanchain[i] <- (<-chanchain[i+1])
	// }
	// chanchain[count-1] <- 10
	// <-chanchain[0]
	// elapse = time.Since(startTime)
	// fmt.Println("chain of Chan", count, " Using ", elapse)