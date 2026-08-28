# 0109 string、[]byte、rune

## 目标

能解释为何 `len(s)` 与 `for range s` 次数可能不同；对应 curriculum 验收。

## 写本篇时的基础假设

- 0102 已会：`string` / `byte` / `rune` 是预声明类型；零值 `""`。  
- 缺口：可能当 Java `char` 用下标；`len` 与 `range` 语义未钉死。

## 正文

### 一句话

Go 的 `string` 是 **UTF-8 字节序列**（只读）。`len` 数字节；`range` 走 rune（码点）。`s[i]` / `[]byte` 是字节层；要按「字」改用 `[]rune`。

### 三层对照

| 东西 | 本质 | 像 Java 的… |
|------|------|-------------|
| `string` | 只读字节序列 | `String`（不可变） |
| `[]byte` | 可变字节切片 | `byte[]` |
| `rune` | Unicode 码点（`int32` 别名） | 更接近「一个字符」的码点，不是 UTF-16 `char` |

### `len` vs `range`

```text
s = "Go中文"
len(s) = 8          ← G o + 中(3) + 文(3)
range 次数 = 4      ← G、o、中、文
range i=2 → '中'
range i=5 → '文'    ← i 是该 rune 的起始字节下标
```

教学小测：`"Hi你好"` → `len=8`，`range=4`（用户已答对）。

### 下标与互转

- `s[i]`：第 i 个 **byte**，不是第 i 个字。`"a中"` 的 `s[1]` 是 `'中'` 的某个字节。  
- `string` **不可变**：`s[0] = 'X'` → **编译失败**。  
- `[]byte(s)` / `string(b)`：**拷贝**一整份。改 `[]byte` 不影响原 `string`。  
- 对中文只改某个字节 → UTF-8 断了，`range` 可能看到 `�`。  
- 按「字」改：`r := []rune(s); r[0] = 'X'; string(r)` → 正经替换第一个码点。

| 写法 | 结果 |
|------|------|
| `s[0] = 'X'` | 编译失败 |
| `[]byte(s)` 再改 `b[0]` | 能跑；中文常弄坏 UTF-8 |
| `[]rune(s)` 再改 `r[0]` | 按码点改字 |

### 对照 Java

| Java | Go |
|------|-----|
| `String.length()` 多数按 UTF-16 code unit | `len(s)` **始终按字节** |
| `charAt` / code point API 分层 | 下标=字节；`range`/`rune`=码点 |
| `getBytes` / `new String(bytes)` 拷贝 | `[]byte(s)` / `string(b)` 同样拷贝 |
| `String` 不可变 | `string` 不可变 |

### 与前置 / 引出

- **0102：** `byte`/`rune`/`string` 类型与零值。  
- **0301：** `io.Reader` 等按字节流组合，心智从本课字节层接上。

### Demo（助手已跑过）

| 路径 | 说明 |
|------|------|
| `demo/` | len/range、下标、坏 UTF-8、`[]rune` 改字 |

## 最小例子

见正文；教学澄清：用户曾不清楚 A/B/C——钉死「不可变 / 字节改坏 / rune 改字」。

## 验收题（已通过）

1. `len` = 字节数；`range s` = rune（码点）次数；含多字节字符时两者常不同。

## 学完后 foundation 应如何改

- **已会：** `len` 字节 / `range` rune；下标与 `[]byte` 是字节层；按字改用 `[]rune`。  
- **缺口去掉：** 「string / []byte / rune」当前条。  
- **仍开放：** 深挖 `unicode/utf8`、字符串 builder 等用到再补。

## 下一步

**0110** struct 与 method。
