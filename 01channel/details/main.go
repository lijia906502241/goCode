package main

import "fmt"

func main() {
	//管道可以声明为只读或者只写
	// 默认情况下，管道是双向的
	//声明为只写
	var intChan chan<- int
	intChan = make(chan int, 3)
	intChan <- 20
	fmt.Println("intChan=", intChan)
	//声明为只读
	var intChan1 <-chan int
	num := <-intChan1
	fmt.Println("num", num)
}
