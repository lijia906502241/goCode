package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func (s Monster) Print() {
	fmt.Println("start----")
	fmt.Println(s)
	fmt.Println("end----")
}
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d:tag为=%v\n", i, tagVal)
		}
	}
	//获取结构体方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods \n", numOfMethod)
	val.Method(1).Call(nil) //获取到第二个方法，调用这个  获取第一个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params) // 传入得参数是 []reflect.Value 返回[]reflect.Value
	fmt.Println("res=", res[0].Int())

}
func main() {
	a := Monster{
		Name:  "黄鼠狼",
		Age:   200,
		Score: 100,
		Sex:   "雌",
	}
	TestStruct(a)
}
