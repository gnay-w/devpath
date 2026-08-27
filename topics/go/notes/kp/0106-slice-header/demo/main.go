package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	a := s[:2]
	b := s[1:3]
	fmt.Printf("改前: s=%v a=%v b=%v\n", s, a, b)
	fmt.Printf("     len/cap: s=%d/%d a=%d/%d b=%d/%d\n",
		len(s), cap(s), len(a), cap(a), len(b), cap(b))

	b[0] = 99
	fmt.Printf("改后: s=%v a=%v b=%v\n", s, a, b)

	// 下标相对当前标签：b 的 [0:1] 不是「原数组下标 0」
	c := b[0:1]
	fmt.Printf("c = b[0:1] → %v (相对 b，不是相对最初的 s)\n", c)
}
