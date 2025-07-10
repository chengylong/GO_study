package main

import "fmt"

// var nums = []int{2, 2, 1}
var nums = []int{4, 1, 2, 1, 2}

//找到切片中只出现一次的数字
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

//回文数
func isPalindrome(x int) bool {
	//负数和以0结尾的正数必然不是回文数
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	reversed := 0
	original := x
	// 反转数字的后半部分
	for original > reversed {
		reversed = reversed*10 + original%10
		original /= 10
	}
	fmt.Println(original)
	fmt.Println(reversed)
	return reversed == original || reversed/10 == original
}

// func main() {
// 	var e = singleNumber(nums)
// 	fmt.Println(e)
// }
