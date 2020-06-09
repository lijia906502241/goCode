package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int //4
	array   [4]int
	head    int
	tail    int
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环状队列情况如下")
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助变量 指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//判断环状为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断环形队列是否空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出环状队列有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
