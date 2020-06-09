package main

import (
	"fmt"
	"reflect"
)

/*
	Type是类型，Kind 是类别，Type和 Kind可能是相同的，也可能是不同的
	var stu Student   stu的Type是 main.Student  Kind是struct
*/
func reflectTest(b interface{}) {
	//通过反射获取传入得变量 type,kind, 值
	// 1.先获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	n := 2 + rVal.Int()
	fmt.Println("n=", n)
	fmt.Printf("rVal=%v type=%T\n", rVal, rVal)
	//将rVal 转换为 interface{}
	iV := rVal.Interface()
	//将interface{} 通过断言转成需要的类型
	num := iV.(int)
	fmt.Println("num=", num)
}

// 对结构体的反射
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iv=%v type=%T\n", iV, iV)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//var num int = 10
	//reflectTest(num)
	stu := Student{
		Name: "lijia",
		Age:  26,
	}
	reflectTest02(stu)
}
