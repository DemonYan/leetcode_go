/*
 * @Author: ruoru
 * @Date: 2020-06-10 14:04:50
 * @LastEditors: Demon
 * @LastEditTime: 2020-06-10 23:25:42
 * @Description: https://leetcode-cn.com/problems/subsets/
 */

 package main

 import (
	 "container/list"
	 "fmt"
	 "math"
)

func backtracking(route string, choices []int, res *list.List) {
	for i := 0; i < len(choices); i++ {
		backtracking(route + string(choices[i]), choices[i+1:], res)
	}
}

func subsets(nums []int) [][]int {
	//数组长度为 2^len(nums) 次方个
	/*
		这里可以用二进制位来类比，比如 [1,2,3]
		其实就对应一个三位的二进制 000
		位数上为1，表示该位数值出现，比如 001 对应的不是 [3]
		所以总共有 8  种情况
		这也是进制位运算可以得到结果的一种解法
	*/
	if nums == nil || len(nums) <= 0 {
		return [][]int{};
	}

	numsSize := len(nums)
	arr := [1][math.Exp2(numsSize)]int{};
	return arr;
}


func main() {
	l := list.New()
	l.PushBack("first")
	l.PushBack("Second")
	test(l)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	
	a := []int
}