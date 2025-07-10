package main

import "fmt"

// var nums = []int{2, 2, 1}
var nums = []int{4, 1, 2, 1, 2}

func main() {
	var e = singleNumber(nums)
	fmt.Println(e)
}
func singleNumber(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for k, v := range countMap {
		if v == 1 {
			return k
			// fmt.Println("只出现一次的元素", k)
		}
	}
	return 0
}
