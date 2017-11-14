package main
import (
	"fmt"
)


func main(){
	var Mysicle = []int{5,7,8,4,3,9,2}
	sort_insert(Mysicle,len(Mysicle))
	fmt.Println(Mysicle)

}


//插入排序
func sort_insert(elems []int,n int){
	for j := 1; j < n;j++{
		key := elems[j]     //要插入的元素
		i := j -1
		for ; i >= 0 && elems[i] > elems[i+1];i-- {    //i会在for中循环减小到不满住条件
			elems[i],elems[i+1] = elems[i+1],elems[i]
		}
		elems[i+1] = key    //key的值刚好比elems[i]大,所以需要i+1,
	}
}