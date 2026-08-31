package main

import "fmt"

// —— 隐式满足 ——

type Writer interface {
	Write([]byte) (int, error)
}

type Buffer struct{}

func (b Buffer) Write(p []byte) (int, error) {
	return len(p), nil
}

func save(w Writer) { fmt.Println("save ok, wrote", len("hi"), "via Writer") }

// —— 方法集：指针接收者 ——

type Closer interface {
	Close() error
}

type File struct{ closed bool }

func (f *File) Close() error {
	f.closed = true
	return nil
}

func shutdown(c Closer) { _ = c.Close(); fmt.Println("shutdown ok") }

// —— 值方法 vs 指针方法进 interface ——

type T struct{ N int }

func (t T) ByValue() { fmt.Println("ByValue N=", t.N) }
func (t *T) ByPointer() {
	t.N++
	fmt.Println("ByPointer N=", t.N)
}

type HasValue interface{ ByValue() }
type HasPtr interface{ ByPointer() }

func main() {
	fmt.Println("=== 1) 隐式满足（无 implements）===")
	var b Buffer
	save(b)

	fmt.Println("=== 2) 方法集：指针方法只有 *T ===")
	var f File
	// shutdown(f) // File does not implement Closer (Close has pointer receiver)
	shutdown(&f)
	fmt.Println("f.closed", f.closed)

	fmt.Println("=== 3) 值进「模拟盒子」改不到外面 ===")
	var x T
	x.N = 10
	copyInBox := x
	(&copyInBox).ByPointer()
	fmt.Println("外面 x.N", x.N, "盒子 copyInBox.N", copyInBox.N)

	fmt.Println("=== 4) 正确：接口装 *T ===")
	var p HasPtr = &x
	p.ByPointer()
	fmt.Println("外面 x.N", x.N)

	fmt.Println("=== 5) 值方法：T 与 *T 都能当 HasValue ===")
	var v HasValue
	v = x
	v.ByValue()
	v = &x
	v.ByValue()
}
