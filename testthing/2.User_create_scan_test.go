package main

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	var (
		userName string
		email    string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)
	fmt.Println(userName, email)
	//	scan输入有问题 无法解决
}
