package main

import (
	"fmt"
)
func main(){

	select_sort([]int{1,5,3,7,2,6,9,8,4},9)
}

func select_sort(nums []int, n int){

	for i := 0; i < n-1; i++ {
		var imin = i
		for j := i+1; j < n; j++ {
			if nums[imin] > nums[j]{
				nums[imin],nums[j] = nums[j],nums[imin]
			}
		}
	}
	fmt.Println(nums)
}

























func dtob(num int)[]int{
	var nums []int
	var  m int
	for num > 0 {
		m = num % 2
		num = num /2
		nums  = append(nums,m)
	}
	reserve(nums)
	return nums
}



func btod(nums []int)int{
	length := len(nums)
	var total int = 0
	var i uint
	for i = 0; i < uint(length) ; i++  {
		total += nums[i] << i
	}
	return total
}


func reserve(nums []int){
	length := len(nums)
	var head = 0
	var rear = length - 1
	for head < rear  {
		nums[head],nums[rear] = nums[rear],nums[head]
		 head++
		 rear--
	}
}


func search(nums []int,val int)int{
	length := len(nums)
	var i int
	for i = 1; i <= length; i++  {
		if val == nums[i-1] {
			return i
		}
	}
	return NULL
}


func insert_sort(nums []int,n int)[]int{
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(nums[1])
		}
	}()
	for i := 2; i <= n;i++{
		for j := i - 1; j > 0 && nums[j] < nums[j-1]  ; j-- {
			nums[j],nums[j-1] = nums[j-1],nums[j]
			fmt.Println(nums[j],nums[j-1])
		}

	}
	return nums
}


