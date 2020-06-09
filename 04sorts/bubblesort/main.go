package main

import "fmt"

/*   冒泡排序
1.先完成能够将最大的数放到最后，把第二大的数放在倒数第二个位置 ........
2.一共会经过arr.length-1次的轮数比较，每一轮将会确定一个数的位置
3. 当发现前面的一个数比后面的一个数大时，就进行了交换
*/
func BubbleSort(arr *[8]int) {
	fmt.Println("排序前arr=", (*arr))
	temp := 0
	for j := 0; j < len(arr)-1; j++ { //循环几轮
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
	fmt.Println("排序后", arr)
}
func main() {
	arr := [8]int{24, 69, 80, 57, 13, 2, 5, 100}
	BubbleSort(&arr)
}
