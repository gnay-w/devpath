package main

import "fmt"

type Engine struct {
	Power int
	Name  string
}

func (e Engine) Start() { fmt.Println("engine", e.Name, e.Power) }

type GPS struct {
	Name string
}

func (g GPS) Start() { fmt.Println("gps", g.Name) }

// CarA：具名字段（普通组合）
type CarA struct {
	Engine Engine
	Brand  string
}

// CarB：嵌入
type CarB struct {
	Engine
	Brand string
}

// Dog：外层同名方法盖住嵌入
type Animal struct{}

func (a Animal) Speak() { fmt.Println("animal") }

type Dog struct {
	Animal
}

func (d Dog) Speak() { fmt.Println("woof") }

// Clash：多嵌入同名 → 须全路径
type Clash struct {
	Engine
	GPS
	Name string
}

func main() {
	// 1) 普通组合 vs 嵌入提升
	a := CarA{Engine: Engine{Power: 150, Name: "V6"}, Brand: "A"}
	b := CarB{Engine: Engine{Power: 150, Name: "V6"}, Brand: "B"}
	fmt.Println("=== 普通组合 ===")
	a.Engine.Start()
	fmt.Println("power", a.Engine.Power)
	fmt.Println("=== 嵌入（可少写一层）===")
	b.Start()
	fmt.Println("power", b.Power, "brand", b.Brand)
	fmt.Println("full path", b.Engine.Power)

	// 2) 外层优先
	fmt.Println("=== 外层 Speak ===")
	Dog{}.Speak()
	Dog{}.Animal.Speak()

	// 3) 同层撞名
	c := Clash{
		Engine: Engine{Name: "V6", Power: 200},
		GPS:    GPS{Name: "Nav"},
		Name:   "我的车",
	}
	fmt.Println("=== 撞名 ===")
	fmt.Println("c.Name", c.Name)
	fmt.Println("c.Engine.Name", c.Engine.Name)
	c.Engine.Start()
	c.GPS.Start()
	// c.Start() // ambiguous selector
}
