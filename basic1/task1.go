package main

import (
	"fmt"
	"math/big"
	"sort"
)

// var nums = []int{2, 2, 1}
var nums = []int{4, 1, 2, 1, 2}

// 1.找到切片中只出现一次的数字
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

// 2.回文数
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

// 3.有效的括号
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

// 4.最长公共前缀
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

/*
5.加一，思路是依次取出数组中的数字作为一个十进制数。对十进制数进行+1操作后，把十进制数转换回数组
*/
func plusOne(digits []int) []int {
	fmt.Println(digits)

	//数组按照顺序转换为int类型整数
	num := new(big.Int)
	for _, v := range digits {
		num = num.Mul(num, big.NewInt(10)).Add(num, big.NewInt(int64(v)))
	}
	//加1
	num.Add(num, big.NewInt(1))
	fmt.Println(num)

	//转换回数组，对10取余，拼接后反转
	result := []int{}
	for num.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		num.DivMod(num, big.NewInt(10), mod) // tmp = tmp / 10, mod = 余数
		result = append(result, int(mod.Int64()))
	}
	//反转数组
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	fmt.Println(result)
	return result
}

// 6.删除有序数组中的重复项
func deleteDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//记录一共还剩几位数
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(nums)
	return k
}

// 两数之和
// func twoSum(nums []int, target int) []int {
// 	//相加操作，在和target 判断是否相等，

// }

//合并区间
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 输入：intervals = [[1,6],[2,4]]
// 输出：[[1,6]]
// 解释：区间 [1,6] 和 [2,4] 可被视为重叠区间。

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	//多条件排序（先按第一个元素，再按第二个元素）
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("排序后:", intervals)
	// 双探针遍历，比较合并后集合中的右节点和要合并的集合中的左节点的大小，做处理
	// 如果右<左，直接加入新的元素，如果右>左，需要合并数组，在判断右元素和新数组右
	// 元素的大小，合并后的数组右元素的值取他们中较大值
	mergInts := [][]int{}
	for i, _ := range intervals {
		//要加入的区间数组的左右元素
		left_e := intervals[i][0]
		rig_e := intervals[i][1]
		merg_len := len(mergInts)
		if i == 0 || (mergInts[merg_len-1])[1] < left_e {
			mergInts = append(mergInts, intervals[i])
		} else {
			//合并区间，确定右边界的具体值
			total_rig := max(mergInts[merg_len-1][1], rig_e)
			mergInts[merg_len-1][1] = total_rig
			fmt.Println("右边界值：", total_rig)
			fmt.Println("mergInts：", mergInts)
		}
	}
	return mergInts
}

func main() {
	// str := "{(())}[()]"
	// bol := validParentheses(str)
	// fmt.Println(bol) "cir","car"
	// strs := []string{"aaa", "aa", "aaa"}

	// strs := []string{"flower", "flow", "fliwht"}
	// ints := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	// ints := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// int := deleteDuplicates(ints)
	ints := [][]int{{1, 6}, {2, 4}, {3, 9}, {9, 12}}
	re_ints := merge(ints)
	fmt.Println(re_ints)

}
