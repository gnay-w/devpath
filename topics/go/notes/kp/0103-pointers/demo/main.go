package main

import "fmt"

type User struct{ Name string }

func renameByValue(u User) {
	u.Name = "Bob"
}

func renameByPtr(u *User) {
	u.Name = "Bob" // 字段自动解引用，等价 (*u).Name
}

func rebind(p *int) {
	x := 100
	p = &x // 只改函数里的指针拷贝
}

func mutate(p *int) {
	*p = 100
}

func main() {
	a := 42
	b := a
	b = 99
	fmt.Printf("值拷贝: a=%d b=%d\n", a, b)

	x := 42
	p := &x
	*p = 99
	fmt.Printf("指针共享: x=%d *p=%d\n", x, *p)

	q := p
	*q = 7
	fmt.Printf("拷贝指针后: x=%d *p=%d *q=%d\n", x, *p, *q)

	u1 := User{Name: "Alice"}
	renameByValue(u1)
	fmt.Printf("传值改 struct: %q\n", u1.Name)

	u2 := User{Name: "Alice"}
	renameByPtr(&u2)
	fmt.Printf("传指针改 struct: %q\n", u2.Name)

	var np *User
	fmt.Printf("指针零值: np==nil → %v\n", np == nil)

	n := 1
	rp := &n
	rebind(rp)
	fmt.Printf("rebind 后: *rp=%d n=%d\n", *rp, n)
	mutate(rp)
	fmt.Printf("mutate 后: *rp=%d n=%d\n", *rp, n)
}
