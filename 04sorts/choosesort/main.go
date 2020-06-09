package main

import "fmt"

/*选择排序*/
func ChooseSort(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		//假设arr[j]是最小值
		min := arr[j]
		minIndex := j
		for i := j + 1; i < len(arr); i++ { // j+1 从第几个开始比较
			// 循环遍历找出最小值，并记下来
			if min > arr[i] {
				min = arr[i]
				minIndex = i
			}
		}
		//交换
		if minIndex != j {
			arr[j], arr[minIndex] = arr[minIndex], arr[j]
		}
		fmt.Printf("第%d排序=%d\n", j, arr)
	}
}
func main() {
	arr := [5]int{2, 5, 4, 7, 1}
	ChooseSort(&arr)
}
