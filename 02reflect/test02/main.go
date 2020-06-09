package main

import (
	"fmt"
	"reflect"
)

/*
	通过反射来修改变量，注意当使用setXXX方法设置需要通过对应的指针类型来完成，这样
	才能改变传入得变量值，同时需要使用到 reflect.Value.Elem() 方法
*/
func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type=%T\n", val)
	val.Elem().SetInt(110)
	fmt.Printf("val=%T\n", val)
}
func main() {
	var num int = 20
	testInt(&num)
	fmt.Println("num=", num)
}
