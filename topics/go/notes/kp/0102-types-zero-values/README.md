# 0102 类型系统基础与零值

## 目标

能列出常用类型零值，并解释为何「零值可用」；对应 curriculum 验收。

## 写本篇时的基础假设

- 写过基本类型；0101 已会包/导出。  
- 零值直觉可能偏 Java `null`；教学中先补了预声明类型清单，再进零值与 map/struct 边界。

## 正文

### 预声明基本类型（日常先熟这批）

`bool`、`string`、`int`/`int64`、`byte`（≡`uint8`）、`rune`（≡`int32`）、`float64`。  
另有定宽整型、`uint`、`complex*` 等；复合类型（slice/map/struct/指针等）另论，但同样有零值。

**对照 Java：** 无包装类型；无 `char`（用 `byte`/`rune`）；`int` 宽度平台相关。

### 零值：声明即可用

每个类型有确定零值；`var x T` 后即可使用，不必先 `null`。

| 类型 | 零值 |
|------|------|
| 数值 | `0` / `0.0` |
| `bool` | `false` |
| `string` | `""`（不是 null） |
| 指针 / slice / map / chan / func / interface | `nil` |
| `struct` | 每个字段各自的零值 |
| 数组 `[n]T` | `n` 个元素皆为零值（**不是** nil） |

指针打印为 `<nil>`，不是整数 `0`。

### 「零值可用」的含义与边界

- **含义：** 许多零值就是合法的「空」（空串、0、`len==0`），不是异常状态；可减少「未初始化 / null」判断。  
- **边界：** 不等于所有操作都安全——nil map **不能写**；nil 指针 **不能解引用**。

### nil map：能读/`len`，不能写

规格保证：nil map 当只读空 map。

- `m == nil` → true；`len(m)` → 0；读缺键 → 值类型零值  
- `m["k"] = v` → panic：`assignment to entry in nil map`  
- 要写：先 `make(map[K]V)` 或字面量 `map[K]V{}`

**为何：** 读/`len` 不需要底层表；写必须有表。设计上方便调用方少判空。

### struct：值字段 vs 指针字段

```go
type User struct {
	Home Addr   // 值：嵌套铺满零值，可直接 .City
	Meta *Addr  // 指针：零值是 nil，解引用会炸
}
```

- 值嵌套：内部字段继续取零值  
- 地址/句柄（指针、map、slice…）：零值停在 `nil`，不会自动 `make`/`new`

`var u User` 已是可用值；不必 `new` 才能碰字段。`new(User)` / `&User{}` 只是得到指向全零内容的指针。

### 切片 vs 数组（易混）

- `var s []int` → nil slice，`len=0`  
- `var a [3]int` → `[0,0,0]`，固定长，**不是** nil  

细节 → 0105。

### Demo（助手已跑过）

| 路径 | 说明 |
|------|------|
| `demo/` | 基本零值、struct 值/指针字段、slice vs array、`make` 后写 map |
| `demo/nilmap-panic/` | nil map 写入 → panic |

## 最小例子

见 `demo/`；核心对照：

```go
var m map[string]int // 可读、可 len；不可写
var u User           // 字段已是零值，可直接用（指针字段除外）
var a [3]int         // [0,0,0]，不是 nil
```

## 验收题（已通过）

1. 常用零值表（含指针是 `nil` 不是 `0`；string 是 `""`）。  
2. 「零值可用」= 合法空状态、少判空；≠ 写 nil map / 解引用 nil 也安全。  
3. 澄清：`[]int` 零值 nil；`[3]int` 零值三格全 0。  
4. struct：值嵌套铺零值；指针字段零值 `nil`。

## 学完后 foundation 应如何改

- **已会：** 预声明常用类型；零值表；零值可用的含义与边界；nil map 读写差；struct 值/指针字段。  
- **缺口去掉：** 「类型与零值系统化」。  
- **仍开放：** 数组 vs 切片细节 → 0105；指针语义 → 0103。

## 下一步

**0103** 指针基础。
