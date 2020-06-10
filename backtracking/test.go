/*
 * @Author: Demon
 * @Date: 2020-06-10 23:04:13
 * @LastEditors: Demon
 * @LastEditTime: 2020-06-10 23:23:40
 * @Description: 
 */

 package main

 import (
	 "fmt"
	 "math"
 )

 func main()  {
	 var size = int(math.Exp2(3));
	 fmt.Println(size)
	 s := 1;
	 fmt.Println(s)
	 for i:=1; i<size; i++ {
		 fmt.Println(i)
	 }
	 var arr [10]int;
	 for i:=0; i<10; i++ {
		 arr[i] = i;
	 }
	 fmt.Println(arr[1:9])
	// var arr [size][1]int;
	//  arr[0][1] = 3;
	//  fmt.Println(arr)
 }
