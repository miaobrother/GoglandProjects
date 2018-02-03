package main

import ("fmt"
		"sort"

)

func basic_sort(a []int)  {
	sort.Ints(a)
}
func bubble_sort(a []int)  {
	for i:= len(a) -1;i>0; i --{
		for j :=0;j<i;j++{
			if a[j] > a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
}


func select_sort(a []int){
	for i := 0;i<len(a);i++{

		for j := i;j<len(a);j++{
			if a[j] <a[i]{
				a[j],a[i] = a[i],a[j]
			}
		}
	}
}

func inser_sort(a []int)  {
	for i :=1 ;i<len(a);i++{
		for j :=i;j>0;j--{
			if a[j] < a[j-1]{
				a[j],a[j-1] = a[j-1],a[j]
			}
		}
	}
}
func partion(a [] int,left,right int) int  {
	var i = left
	var j = right
	for i <j{
		for j>i && a[j] >a[left]{
			j--
		}
		a[j],a[left] = a[left],a[j]
		for i <j && a[i]<a[left]{
			i++
		}
		a[left],a[i] = a[i],a[left]

	}
	return i
}

func fast_sort(a [] int,left,right int)  {
	if left >= right {
		return
	}
	mid := partion(a,left,right)
	fast_sort(a,left,mid-1)
	fast_sort(a,mid+1,right)
}
func main()  {
	a :=[]int{38,3,5,2,99,88,90}
	fast_sort(a,0,len(a)-1)
	//inser_sort(a)//插入排序算法
	//basic_sort(a)//系统自带排序
	//select_sort(a)//选择排序算法
	//bubble_sort(a)//冒泡排序算法
	fmt.Println(a)
}