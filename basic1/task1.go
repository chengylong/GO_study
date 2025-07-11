package main

import (
	"fmt"
)

// var nums = []int{2, 2, 1}
var nums = []int{4, 1, 2, 1, 2}

// 找到切片中只出现一次的数字
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

// 回文数
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

// 有效的括号
func validParentheses(s string) bool {
	// str = "()[]{}"
	//字符串必须是长度大于0且长度为偶数
	len1 := len(s)
	if len1 == 0 || len1%2 != 0 {
		return false
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len1; i++ {
		//取到右括号
		if m[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			//匹配到栈最上面的左括号，出栈
			stack = stack[:len(stack)-1]
			//左括号入站
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 最长公共前缀
// 其实就是去一个字符串数组中，所有元素前缀相同的前缀·
func publicPrefix(strs []string) string {
	// strs = []string{"flower", "flow", "flight"}
	len1 := len(strs)
	if len1 == 0 {
		return ""
	}
	if len1 == 1 {
		return strs[0]
	}
	preStr := strs[0]
	//比较字符串元素
	for i := 1; i <= len(strs)-1; i++ {
		//需要取前缀再比较，不可以用第一个元素一直比较，否则会导致取不到公共前缀，可能是第一个和最后一个前缀
		preStr = dealPrefix(preStr, strs[i])
		fmt.Println("preStr:", preStr)
		//未匹配上直接中断
		if len(preStr) == 0 {
			return ""
		}
	}
	return preStr
}

// 获取公共前缀，没有什么最长的概念，就是获取两个字符串公共前缀罢了
func dealPrefix(str1 string, str2 string) string {
	//循环次数为较短的字符串长度
	lenth := min(len(str1), len(str2))
	//记录前n位相同字符
	index := 0
	for i := 0; i < lenth; i++ {
		if str1[i] == str2[i] {
			index++
		} else {
			break
		}
	}
	return str1[:index]
}

func main() {
	// str := "{(())}[()]"
	// bol := validParentheses(str)
	// fmt.Println(bol) "cir","car"
	// strs := []string{"aaa", "aa", "aaa"}

	strs := []string{"flower", "flow", "fliwht"}
	preStr := publicPrefix(strs)
	fmt.Println(preStr)

}
