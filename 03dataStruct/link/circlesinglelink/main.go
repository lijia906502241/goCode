package main

import (
	"fmt"
)

type Heros struct {
	id       int
	name     string
	nickname string
	next     *Heros
}

func InsertHero(head *Heros, newHero *Heros) {
	if head.next == nil {
		head.id = newHero.id
		head.name = newHero.name
		head.nickname = newHero.nickname
		head.next = head //构成一个环状
		fmt.Println("首个英雄加入的是", newHero)
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加到链表
	temp.next = newHero
	newHero.next = head
}

//输出这个环状链表
func ListCircleLink(head *Heros) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]->", temp.id, temp.name, temp.nickname)
		if temp.next == head {
			fmt.Println("到头了...")
			break
		}
		temp = temp.next
	}
}

//删除英雄
func DelHeros(head *Heros, id int) *Heros {
	temp := head
	help := head
	if temp.next == nil { //空链表
		fmt.Println("这是一个空的")
		return head
	}
	if temp.next == head {
		if temp.id == id {
			temp.next = nil
		}
		return head
	}
	for {
		if help.next == head {
			break
		}
		help = help.next
	}
	flag := true
	for {
		if temp.next == head {
			break
		}
		if temp.id == id {
			if temp == head {
				head = head.next
			}
			help.next = temp.next
			flag = false
			break
		}
		temp = temp.next
		help = help.next
	}
	if flag {
		if temp.id == id {
			help.next = temp.next
		} else {
			fmt.Printf("sorry，没有对应的id=%d", id)
		}
	}
	return head
}
func main() {
	head := &Heros{}
	hero := &Heros{
		id:       1,
		name:     "疾风剑豪",
		nickname: "亚索",
	}
	hero2 := &Heros{
		id:       2,
		name:     "放逐之刃",
		nickname: "瑞雯",
	}
	hero3 := &Heros{
		id:       3,
		name:     "盲僧",
		nickname: "瞎子",
	}
	InsertHero(head, hero)
	InsertHero(head, hero2)
	InsertHero(head, hero3)
	ListCircleLink(head)
	//DelHeros(head,1)
	ListCircleLink(head)
}
