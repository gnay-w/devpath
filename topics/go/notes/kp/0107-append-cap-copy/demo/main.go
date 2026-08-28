package main

import "fmt"

func main() {
	// 1) 有空位：不换底层，旧切片仍共享
	s1 := make([]int, 2, 4)
	old1 := s1
	s1 = append(s1, 10)
	s1[0] = 99
	fmt.Printf("有空位: s1=%v old1=%v  (共享)\n", s1, old1)

	// 2) 满了：换底层，旧切片分手
	s2 := make([]int, 2, 2)
	old2 := s2
	s2 = append(s2, 10)
	s2[0] = 99
	fmt.Printf("满了扩容: s2=%v old2=%v  (不共享)\n", s2, old2)

	// 3) 验收对照：满 → a 不受影响；有空位 → a 跟着变
	s := []int{1, 2, 3} // len=cap=3
	a := s[:2]
	s = append(s, 4)
	s[0] = 99
	fmt.Printf("满了再改: a=%v\n", a) // [1 2]

	s3 := make([]int, 3, 5)
	s3[0], s3[1], s3[2] = 1, 2, 3
	a3 := s3[:2]
	s3 = append(s3, 4)
	s3[0] = 99
	fmt.Printf("有空位再改: a3=%v\n", a3) // [99 2]

	// 4) copy：稳妥独立副本
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src)
	src[0] = 99
	fmt.Printf("copy %d 个后改 src: src=%v dst=%v\n", n, src, dst)
}
