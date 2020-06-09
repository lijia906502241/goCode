package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() { //  这里使用defer+recover异常处理模式
	// 还有自定义错误    被调用者：errors.New("----")  调用者：panic(err)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //err
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
