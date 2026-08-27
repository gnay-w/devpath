package main

import "fmt"

type Addr struct {
	City string
}

type User struct {
	Name string
	Age  int
	Home Addr
	Meta *Addr
}

func main() {
	var i int
	var s string
	var b bool
	var p *int
	fmt.Printf("basics: i=%v s=%q b=%v p=%v (p==nil? %v)\n", i, s, b, p, p == nil)

	var u User
	fmt.Printf("struct: %#v\n", u)
	fmt.Println("Home.City OK:", u.Home.City)
	fmt.Println("Meta == nil?", u.Meta == nil)
	u.Meta = &Addr{City: "Shanghai"}
	fmt.Println("Meta.City after assign:", u.Meta.City)

	var sl []int
	var a [3]int
	fmt.Printf("slice: %#v nil? %v len=%d\n", sl, sl == nil, len(sl))
	fmt.Printf("array: %#v len=%d\n", a, len(a))

	m := make(map[string]int)
	m["a"] = 1
	fmt.Println("make map then write:", m)
}
