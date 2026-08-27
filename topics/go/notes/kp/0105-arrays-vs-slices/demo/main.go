package main

import "fmt"

func bumpArray(x [3]int) {
	x[0] = 99
	fmt.Println("  函数里数组:", x)
}

func bumpSlice(x []int) {
	x[0] = 99
	fmt.Println("  函数里切片:", x)
}

func main() {
	var a [3]int
	var s []int
	fmt.Printf("零值: a=%v (%T)  s=%v (%T, nil=%v)\n", a, a, s, s, s == nil)

	// s = a  // 编译失败：cannot use a ([3]int) as []int
	s = a[:]
	fmt.Printf("s = a[:] → %v (%T)\n", s, s)

	arr := [3]int{1, 2, 3}
	fmt.Println("\n=== 传数组（整盒拷贝）===")
	fmt.Println("调用前:", arr)
	bumpArray(arr)
	fmt.Println("调用后:", arr)

	sl := []int{1, 2, 3}
	fmt.Println("\n=== 传切片（拷标签，共享底层）===")
	fmt.Println("调用前:", sl)
	bumpSlice(sl)
	fmt.Println("调用后:", sl)

	fmt.Println("\n=== 赋值同理 ===")
	b := arr
	b[0] = 7
	fmt.Println("数组赋值后: arr=", arr, "b=", b)
	t := sl
	t[0] = 7
	fmt.Println("切片赋值后: sl=", sl, "t=", t)
}
