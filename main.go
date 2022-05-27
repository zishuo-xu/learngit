package main

import (
	"fmt"
)

//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名结构体
}

func main() {
	var user2 User
	user2.Name = "pprof"
	user2.Gender = "女"
	user2.Address.Province = "黑龙江"   //通过匿名结构体.字段名访问
	user2.City = "哈尔滨"               //直接访问匿名结构体的字段名
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
}
