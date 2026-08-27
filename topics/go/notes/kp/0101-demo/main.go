package main

import (
	"fmt"

	"devpath.local/go0101/user"
)

func main() {
	u := user.NewUser("gnay")
	fmt.Println("via Name():", u.Name())

	// 故意访问未导出字段 —— 应编译失败
	fmt.Println(u.name)
}
