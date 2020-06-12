/*
 * @Author: ruoru
 * @Date: 2020-06-10 14:04:50
 * @LastEditors: ruoru
 * @LastEditTime: 2020-06-12 17:38:01
 * @Description: https://leetcode-cn.com/problems/subsets/
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func backtracking(route string, choices []int, cmap map[string]int) {
	if len(choices) == 0 {
		cmap[route] = 1
		return
	}

	for i := 0; i < len(choices); i++ {
		backtracking(route+fmt.Sprint(choices[i])+",", choices[i+1:], cmap)
		backtracking(route, choices[i+1:], cmap)
	}
}

func subsets(nums []int) [][]int {
	arr := make([][]int, 0)
	if len(nums) == 0 {
		return arr
	}
	cmap := make(map[string]int)
	backtracking("", nums, cmap)
	for k, _ := range cmap {
		if len(k) == 0 {
			arr = append(arr, make([]int, 0))
			continue
		}
		var tmp = strings.Split(k[0:len(k)-1], ",")
		var tmpInt = make([]int, 0)
		for _, v := range tmp {
			iv, _ := strconv.Atoi(v)
			tmpInt = append(tmpInt, iv)
		}
		arr = append(arr, tmpInt)
	}
	return arr
}

func main() {
	// arr := make([][]int, 0, 10)
	// t := make([]int, 0, 10)
	// for i := 0; i < 12; i++ {
	// 	t = append(t, i)
	// 	fmt.Printf("%p \n", t)
	// }
	// fmt.Println(t)
	// arr = append(arr, t)
	// fmt.Println(t)
	// fmt.Println(arr)

	// fmt.Println(strings.Split("abc", ""))
	// var arr []int
	// var arr2 = append(arr, 1)
	// fmt.Println(arr)
	// fmt.Println(arr2)
	// fmt.Println(arr)
	// var arr = []int{1, 2, 3}
	// var arr2 = arr[0:]
	// fmt.Printf("%p \n", arr)
	// fmt.Printf("%p \n", arr2)
	// copy(arr, arr[0:])
	// fmt.Printf("%p \n", arr)
	// // arr2[1] = 3
	// fmt.Println(arr)
	// reverse(arr)
	// fmt.Println(arr)
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 5}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))

	// var a = []int{1, 2, 3}
	// b, _ := json.Marshal(a)
	// fmt.Println(string(b))
	// var m = make(map[string]int)
	// m[string(b)] = 1
	// fmt.Println(m)
	// fmt.Println(fmt.Sprint(a))
}
