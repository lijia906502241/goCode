package main

import (
	"fmt"
	"net/http"
	"time"
)

var url = []string{
	"https://www.baidu.com",
	"https://www.google.com",
	"https://www.taobao.com",
}

func main() {
	for _, v := range url {
		//resp , err := http.Head(v)
		c := http.Client{
			Timeout: time.Second * 2,
		}
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err=%v\n", v, err)
			continue
		}
		fmt.Printf("head success! ==> status=%v\n", resp.Status)
	}
}
