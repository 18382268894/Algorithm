/*
   递归实现插入排序
   时间复杂度为O(n^2)
*/


package main

import (
	"fmt"
)

func insert(nums []int, i int){

	for ;i >= 1 && nums[i-1] > nums[i] ; i-- {
		nums[i-1],nums[i] = nums[i],nums[i-1]
	}
}

func insert_sort(nums []int, n int){
	if n > 0{
		n--
		insert_sort(nums,n)
		insert(nums,n)
	}
}

func main(){
	nums := []int{4,2,6,8,9,7,1,82,43,76,54}
	insert_sort(nums,len(nums))
	fmt.Println(nums)
}
