package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个结点
	next     *HeroNode //指向后一个结点
}

//插入数据到链表，在单链表的最后插入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到该链表的最后这个结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp 不断的指向下一个结点
	}
	//将newHeroNode 加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 根据 no的大小插入
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
		fmt.Println("sorry，已经存在no=", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表》。。。")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]=>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//逆序遍历这个双向链表
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表......")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%d,%s,%s]=>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

//双向链表删除一个结点
func DeleteHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for { // 找到要删除结点的no，和temp的下一个结点的no比较
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry,要删除的ID不存在")
	}
}
func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "放逐之刃",
		nickname: "锐雯",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "疾风剑豪",
		nickname: "亚索",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "盲僧",
		nickname: "瞎子",
	}
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	//InsertHeroNode2(head,hero3)
	//InsertHeroNode2(head,hero1)
	//InsertHeroNode2(head,hero2)
	ListHeroNode(head)

	DeleteHeroNode(head, 2)
	ListHeroNode2(head)
}
