package main

import "fmt"

func main() {
	// 1) len = 字节；range = rune
	s := "Go中文"
	fmt.Printf("s=%q len=%d\n", s, len(s))
	n := 0
	for range s {
		n++
	}
	fmt.Printf("range 次数=%d\n", n)

	// 2) 下标是字节；range 的 i 是该 rune 的起始字节下标
	fmt.Printf("s[0]=%q (byte)\n", s[0])
	for i, r := range s {
		fmt.Printf("range i=%d rune=%q U+%04X\n", i, r, r)
	}

	// 3) []byte 拷贝后可改；改中文字节会弄坏 UTF-8
	b := []byte("你好")
	b[0] = 'X'
	fmt.Printf("坏 UTF-8: %q\n", string(b))

	// 4) []rune 按码点改才是「改一个字」
	r := []rune("你好")
	r[0] = 'X'
	fmt.Printf("按字改: %q\n", string(r))

	// string 不可变：s[0] = 'X' 会编译失败（勿写）
}
