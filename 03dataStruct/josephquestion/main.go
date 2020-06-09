package main

import (
	"fmt"
	"strconv"
)

type Boy struct {
	No   int
	Name string
	Next *Boy //指向下一个小孩的指针
}

// num表示小孩个数，  *Boy 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num的值小于1")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No:   i,
			Name: "tom" + strconv.Itoa(i),
		}
		// 分析构成循环链表，需要一个辅助指针 帮忙
		if i == 1 { // 第一个小孩
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构造环状链表
		}
	}
	return first
}

// 显示单向的环形链表
func ShowBoy(first *Boy) int {
	var count int = 1
	if first.Next == nil {
		fmt.Println("链表为空。。。")
		return 0
	}
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d,小孩名字=%s-->\n", curBoy.No, curBoy.Name)
		if curBoy.Next == first {
			break
		}
		//curBoy 移动到下一个
		curBoy = curBoy.Next
		count++
	}
	return count
}

/*编写一个函数，在环形链表中留下最后一个人*/
func PlayGame(first *Boy, startNo int, countNum int, count int) {
	// 1.判断空链表
	if first.Next == nil {
		fmt.Println("链表为空。。。")
		return
	}
	fmt.Println("小孩的总数=", count) //判断起始号不能大于总数
	if startNo > count {
		fmt.Println("开始序号大于总数")
		return
	}
	// 2. 定义一个辅助指针，帮助删除小孩
	tail := first
	// 3. 让 tail执行到环形链表的最后一个小孩，tail在删除小孩时需要使用
	for {
		if tail.Next == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	// 4. 让first移动到 startNo【后面删除小孩，以first为准】
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	// 5. 开始数countNum,然后就删除first指向的小孩
	for {
		//开始数countNum -1 次
		for i := 1; i < countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d,出圈\n", first.No)
		//删除first执行的小孩
		first = first.Next
		tail.Next = first
		// 如果 tail == first ，圈子中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后一个小孩编号为%d 出圈\n", first.No)
}

func main() {
	first := AddBoy(500)
	count := ShowBoy(first)
	PlayGame(first, 8, 31, count)
}
