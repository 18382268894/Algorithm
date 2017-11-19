/*
	归并排序
*/

package main

import (
	"fmt"
	"math"
)

/*
   合并子序列
   nums中有两个已排序的子序列,合并两个子序列,重新写入nums中
   p为第一个子序列第一个元素下标
   q第二个子序列第一个元素的下标,
   r为第二个元素最后一个下标
*/
func merge(nums []int, p int,q int,r int){
	n1 := q - p + 1    //第一个子序列长度
	n2 := r - q    //第二个子序列长度,第一个元素下标为q+1
	slice1 := make([]int,n1+1)  //第一个子序列
	slice2 := make([]int,n2+1)  //第二个子序列
	slice1[n1] = math.MaxInt64  //设置哨兵值,视为int类型的无穷大
	slice2[n2] = math.MaxInt64
	for i := 0; i < n1; i++ {
		slice1[i] = nums[i+p]
	}
	for j := 0; j < n2 ; j++ {
		slice2[j] = nums[j+q+1]
	}
	count1, count2 := 0, 0
	for k := p; k <= r; k++ {
		//将slice1和slice2中较小的数放入nums中,由于哨兵值无限大,所以不会被放入nums中
		if slice1[count1] <= slice2[count2] {
			nums[k] = slice1[count1]
			count1++
		}else {
			nums[k] = slice2[count2]
			count2++
		}
	}
}

/*
	归并排序
    p为排序序列的第一个元素的下标,第二个元素为最后一个元素额下标
*/
func merge_sort(nums []int, p int, r int){
	if   r > p {        //递归出口,当子序列无法再分时,r = p,此时开始依次调回merge()
		q := (r+p)/2    //二分点,用于将序列分为两个子序列
		merge_sort(nums,p,q)   //归并排序左子序列
		merge_sort(nums,q+1,r)  //归并排序右子序列
		merge(nums,p,q,r)      //合并左右子序列
	}
}


func main(){
	nums := []int{21,12,10,5,78,60,23,8,9,105,107,54,33}
	merge_sort(nums,0,len(nums)-1)
	fmt.Println(nums)
}




