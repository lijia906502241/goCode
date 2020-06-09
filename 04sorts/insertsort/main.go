package main

import (
	"fmt"
	"time"
)

/*插入排序*/
func InsertSort(arr *[5]int) {
	for j := 1; j < len(arr); j++ {
		// 完成第一次，给第二个元素找到合适的位置并插入
		insertVal := arr[j]
		insertIndex := j - 1
		//从大到小
		for insertIndex >= 0 && insertVal > arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次排序=>%v\n", j, *arr)
	}
}

//// 完成第一次，给第二个元素找到合适的位置并插入
//insertVal := arr[1]
//insertIndex := 1 - 1
////从大到小
//for insertIndex >= 0 && insertVal >arr[insertIndex] {
//	arr[insertIndex +1] = arr[insertIndex]
//	insertIndex --
//}
////插入
//if insertIndex + 1 != 1{
//	arr[insertIndex +1] = insertVal
//}
//fmt.Println("排序后",*arr)
func main() {
	arr := [5]int{1, 23, 26, 31, 7}
	before := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Println(end - before)
}
