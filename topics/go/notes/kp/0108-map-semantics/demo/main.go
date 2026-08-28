package main

import "fmt"

func main() {
	// 1) 缺键 → 值类型零值（int 是 0，不是 nil）
	m := map[string]int{"a": 1}
	fmt.Printf("m[\"a\"]=%d m[\"b\"]=%d\n", m["a"], m["b"])

	// 2) 逗号 ok：区分「没有」和「有且为零」
	m["a"] = 0
	v1, ok1 := m["a"]
	v2, ok2 := m["missing"]
	fmt.Printf("有键零值: %d %v | 缺键: %d %v\n", v1, ok1, v2, ok2)

	// 3) 值类型决定缺键长什么样
	ms := map[string]string{}
	mp := map[string]*int{}
	fmt.Printf("缺 string 键=%q | 缺 *int 键==nil? %v\n", ms["x"], mp["x"] == nil)

	// 4) 并发写会 fatal（勿默认打开；需要时取消注释）
	// 两个 goroutine 同时 m[k]++ → fatal error: concurrent map writes
	// go run . 会直接退出；recover 救不了。
}
