package main

import "fmt"

type Counter struct{ N int }

func (c Counter) IncValue() { c.N++ }
func (c *Counter) IncPtr()  { c.N++ }

type User struct{ Name string }

func (u User) RenameVal(n string)  { u.Name = n }
func (u *User) RenamePtr(n string) { u.Name = n }

func NewUser(name string) *User {
	return &User{Name: name}
}

func main() {
	// 1) 值接收者改拷贝；指针接收者改原件
	c := Counter{N: 0}
	fmt.Println("start", c.N)
	c.IncValue()
	fmt.Println("after IncValue", c.N)
	c.IncPtr()
	fmt.Println("after IncPtr", c.N)

	// 2) Rename：值 vs 指针
	u := User{Name: "Alice"}
	u.RenameVal("Bob")
	fmt.Println("after RenameVal", u.Name)
	u.RenamePtr("Bob")
	fmt.Println("after RenamePtr", u.Name)

	// 3) 构造约定
	p := NewUser("Carol")
	p.RenamePtr("Dan")
	fmt.Println("NewUser then RenamePtr", p.Name)

	// 4) 语法糖：值变量可调指针方法（可取址时）
	c2 := Counter{N: 10}
	c2.IncPtr() // ≡ (&c2).IncPtr()
	fmt.Println("sugar IncPtr on value", c2.N)
}
