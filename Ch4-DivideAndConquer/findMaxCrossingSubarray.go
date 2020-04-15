package main

import (
	"fmt"
	"math"
)

func findMaxCrossingSubarray(array []int, low, mid, high int) (int, int, int) {

	intMinimum := ^int(^uint(0) >> 1)
	leftSum, rigthSum := intMinimum, intMinimum
	sum := 0
	var maxLeft, maxRight int
	for i := mid; i >= low; i-- {
		sum = sum + array[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum = sum + array[i]
		if sum > rigthSum {
			rigthSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rigthSum
}

func findMaximumSubarray(array []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, array[low]
	}
	mid := int(math.Floor(float64((low + high) / 2)))
	leftLow, leftHigh, leftSum := findMaximumSubarray(array, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(array, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(array, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}
	if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}

func main() {
	example := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	low, high, sum := findMaximumSubarray(example, 0, len(example)-1)
	fmt.Printf("the maximum subarray of the array is from %d to %d, the sum is %d.", low, high, sum)
}
