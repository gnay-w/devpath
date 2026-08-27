# 0101 包、模块可见性与目录

## 目标

能说清：导出规则、`internal/`、import 路径与「import 的是包」；对应 curriculum 验收。

## 写本篇时的基础假设

- 有 Java `public`/`private`/`package` 经验；能读 Go、会 `import`。  
- 对「首字母导出」「import path 怎么拼」含糊；教学中卡在 **import 路径 = module + 目录**，并追问 **import 是不是 package**。

## 正文

### 导出：看首字母

- 大写开头：包外可见（导出）  
- 小写开头：仅**同一 package** 可见  
- 无 `public`/`private`/`protected` 关键字  

常见封装：字段小写 + 导出的 `NewUser` / `Name()`（Go 习惯短名，不必 `Get` 前缀）。

**对照 Java：** 像 private 字段 + public 构造/getter，但靠命名，没有 protected 档；跨包「受控共享」多用 `internal/`。

### `internal/`：编译器强制

路径中含 `internal` 的包，仅其**父目录树之内**可 import；其他 module 会报错：`use of internal package ... not allowed`。

### import 路径 vs package 名

1. **`go.mod` 的 module 名 + 目录 = import path**  
   例：`module devpath.local/go0101` + 目录 `user/` → `import "devpath.local/go0101/user"`  
2. **import 的是包（package）**，不是单个 `.go` 文件，也不是「与 package 无关的目录复制」。  
   path → 定位到目录 → 目录内同一个 `package` 名。  
3. 代码里用的前缀一般是 **package 名**（如 `user.NewUser`），来自文件里的 `package user`，不是随便起的别名（可用别名，以后再说）。

**教学中的困惑与澄清：**

- 易说成「文件名 user」→ 实为**目录名** `user/`（文件常是 `user.go`）。  
- 问「import 一定是 package 吗？」→ **是**：import 引入的是包；path 只是找到它的地址。

### Demo（助手已跑过）

| 路径 | 说明 |
|------|------|
| `0101-demo/` | `NewUser` / `Name()`；访问 `u.name` 会编译失败 |
| `0101-demo/cmd/ok` | 同 module 可 import `internal/store` |
| `0101-outsider/` | 跨 module import `internal` → 被拒绝 |

## 最小例子

见上述 demo；核心 import：

```go
import "devpath.local/go0101/user" // module + 目录 → 包 user
```

## 验收题（已通过）

1. 包外不能 `u.name`、能 `u.Name()` —— 小写不可见 / 大写方法导出。  
2. `internal` —— 编译器强制。  
3. import —— 目录对应的**包**；不是文件名。并确认：import 路径指向一个包。

## 学完后 foundation 应如何改

- **已会：** 首字母导出；`NewXxx`+getter 封装；`internal/` 强制边界；import path = module+目录；import 的是包。  
- **缺口去掉：** 「导出规则可能含糊」。  
- **仍开放：** 多包布局 / `package xxx_test` → 后续 KP。

## 下一步

**0102** 类型系统基础与零值。
