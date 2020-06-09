package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    //表示我们栈最大可以存放数的个数
	Top    int    //表示栈顶，因为栈顶固定，因此我们直接使用Top
	arr    [5]int // 数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error) {
	//判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == 0 {
		fmt.Println("stack full")
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//遍历
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("%d->%d\n", i, this.arr[i])
	}
}
func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, //当为0 时表示空
	}
	stack.Push(11)
	stack.Push(12)
	stack.Push(13)
	stack.Push(14)
	stack.Push(15)
	stack.List()
	val, _ := stack.Pop()
	fmt.Println("出栈->", val)
}
