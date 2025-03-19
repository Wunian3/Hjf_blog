package main

import (
	"blog_server/utils/random"
	"fmt"
)

func main() {
	s := random.RandStr(16)
	fmt.Println(s)
}
