package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

// 编写第一种插入方法，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// temp 辅助结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	// 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("用户ID已存在")
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

// 显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 删除一个结点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id { //找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("sorry,要删除的ID不存在")
	}
}
func main() {
	head := &HeroNode{}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "花荣",
		nickname: "小李广",
	}
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)

	ListHeroNode(head)
	DelHeroNode(head, 3)
	ListHeroNode(head)
}
