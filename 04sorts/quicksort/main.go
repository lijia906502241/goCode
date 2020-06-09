package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
	left表示 数组左边的下标
	right 表示 数组右边的下标
	arr 表示 要排序的数组
*/
func QuickSort(left int, right int, arr *[8000000]int) {
	l := left
	r := right
	//pivot是中轴，支点
	pivot := arr[(left+right)/2]
	// for循环 是将比pivot小的数放到左边，比pivot大的数放到右边
	for l < r {
		//从pivot的左边找到大于等于pivot的值
		for arr[l] < pivot {
			l++
		}
		//从pivot的右边找到小于等于pivot的值
		for arr[r] > pivot {
			r--
		}
		// l >= r 说明本次分解任务完成，break
		if l >= r {
			break
		}
		//交换
		arr[l], arr[r] = arr[r], arr[l]
		//优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	//如果 l == r ,在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}
func main() {
	//测试结果， 800w的数据只需要1s
	var arr [8000000]int
	for i := 0; i < 800; i++ {
		arr[i] = rand.Intn(90000000)
	}
	fmt.Println(arr)
	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Println("快速排序排800w数据用时=", strconv.FormatInt(end-start, 10)+"s")
}
