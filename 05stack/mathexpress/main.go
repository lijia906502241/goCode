package main

import (
	"fmt"
)

// 八皇后问题
func main() {
	var balance = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	queen(balance, 0)
}
func queen(a [8]int, cur int) {
	if cur == len(a) {
		fmt.Print(a)
		fmt.Println()
		return
	}
	for i := 0; i < len(a); i++ {
		a[cur] = i
		flag := true
		for j := 0; j < cur; j++ {
			ab := i - a[j]
			temp := 0
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}
		if flag {
			queen(a, cur+1)
		}
	}
}
