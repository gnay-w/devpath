package main

import (
	"fmt"

	"devpath.local/go0101/user"
)

func main() {
	u := user.NewUser("gnay")
	fmt.Println("via Name():", u.Name())

	// 取消下一行注释再 go run，看编译器怎么说：
	// fmt.Println(u.name)
}
