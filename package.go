package main

import (
	"math"
	"sort"
)

//875
func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	//最小是1
	var left = 1
	var right = piles[len(piles)-1]
	var mid = (left + right) / 2
	for left < right {
		mid = (left + right) / 2
		if checkEat(piles, mid, h) {
			left = mid + 1
		} else {
			//因为出现相等的情况时，就不需要再-1
			right = mid
		}
	}
	return left
}
func checkEat(array []int, speed int, h int) bool {
	var sum = 0
	for _, x := range array {
		sum += int(math.Ceil(float64(x) / float64(speed)))
	}
	return sum > h
}
func main() {
	// piles := []int{30, 11, 23, 4, 20}
	// h := 6
	// fmt.Println(minEatingSpeed(piles, h))

	//fmt.Println(ics4.Sqrt(9))
	//nums := []int{3, 4, 5, 1, 2}
	//nums := []int//{4, 5, 6, 7, 0,, 2}
	//nums := []int{11, 13, 15, 17}
	//nums := []int{2, 1}
	//fmt.Print(FindMin(nums))
}
