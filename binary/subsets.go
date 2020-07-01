/*
 * @Author: ruoru
 * @Date: 2020-07-01 11:52:26
 * @LastEditors: ruoru
 * @LastEditTime: 2020-07-01 12:03:43
 * @Description:
 */
package main

import "fmt"

/**
基本思路：
利用二进制：
比如 [1, 2, 3]，子集的情况应该是 2^3=8 个
分别对应
0：表示000，即所有位对应的数字都不出现
1：表示001，即第1位数出现，则结果对应3，依此类推，直到7

所以总共分两步：
n 为传入数组长度：
1. 外层循环遍历 0 - 2^n，2^n 对应 1 << n
2. 每次与对应位置为1，其他位置为0的元素做 & 运算

*/
func subsets(nums []int) [][]int {
	arr := make([][]int, 0)
	if len(nums) == 0 {
		return arr
	}
	var nlen = len(nums)
	for i := 0; i < 1<<nlen; i++ {
		tmp := make([]int, 0)
		for j := 0; j < nlen; j++ {
			if (1<<j)&i > 0 {
				tmp = append(tmp, nums[j])
			}
		}
		arr = append(arr, tmp)
	}
	return arr
}

func main() {
	fmt.Println(1 << 2)
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 5}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
