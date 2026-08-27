package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println("nil?", m == nil, "len=", len(m), "read=", m["a"])
	fmt.Println("writing to nil map…")
	m["a"] = 1 // panic: assignment to entry in nil map
}
