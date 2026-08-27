package main

import "fmt"

type User struct{ Name string }

// 试探题：返回局部指针在 Go 里安全（x 会逃逸上堆）
func f() (*int, int) {
	x := 1
	p := &x
	return p, x
}

// 只返回值：地址没出去，通常不必逃逸
func stayOnStack() int {
	y := 2
	return y
}

func returnByValue() User {
	u := User{Name: "a"}
	return u
}

func returnByPtr() *User {
	u := User{Name: "b"}
	return &u
}

// 返回 any：接口盒子常带着指向数据的指针 → 内容易上堆
func returnAny() any {
	n := 42
	return n
}

func main() {
	p, v := f()
	fmt.Printf("f: *p=%d v=%d\n", *p, v)
	fmt.Printf("stayOnStack: %d\n", stayOnStack())
	fmt.Printf("returnByValue: %+v\n", returnByValue())
	fmt.Printf("returnByPtr: %+v\n", returnByPtr())
	fmt.Printf("returnAny: %T %v\n", returnAny(), returnAny())
}
