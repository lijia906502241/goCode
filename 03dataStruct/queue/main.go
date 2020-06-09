package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int //数组 -> 模拟队列
	front   int    // 表示指向队列首
	rear    int    // 表示指向队列尾
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//从队列中取出元素
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

//显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	queue := Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入1表示添加数据到队列")
		fmt.Println("2.输入2表示从队列获取数据")
		fmt.Println("3.输入3表示显示队列")
		fmt.Println("4.输入4表示退出队列")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("输入你要的入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok...")
			}
		case "2":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "3":
			queue.ShowQueue()
		case "4":
			os.Exit(0)
		}
	}
}
