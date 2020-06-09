package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	fmt.Println("ok")
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
