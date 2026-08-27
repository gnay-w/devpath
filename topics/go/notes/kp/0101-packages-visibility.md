# 0101 包、模块可见性与目录

## 目标

学完能说清：Go 里「谁能看见谁」、目录和包的关系、以及 `internal/` 是干什么的（对应 curriculum 验收）。

## 写本篇时的基础假设

- **已会：** 能 `import`、能读常见语法；有 Java 包/`public` 经验。  
- **缺口：** 导出规则与目录约定可能含糊；尚未按编号系统打语言底子。

## 正文

### 1. 包是什么

- 一个目录里的 `.go` 文件通常声明**同一个** `package` 名。  
- 包名一般是目录名（小写）；`main` 是特例（可执行入口）。  
- 同一包内：随便互相调用，**无** Java 那种「同类才 private 字段」的额外层——可见性只看**首字母**（见下）。

### 2. 可见性：首字母，不是 public 关键字

| Go | 大致对应 Java | 含义 |
|----|---------------|------|
| `Hello`（大写开头） | `public` | **导出**：其他包可引用 |
| `hello`（小写开头） | 包私有（比 Java default 更干净） | **仅本包**可见 |

适用于：函数、类型、结构体字段、常量、变量、方法名。

```go
package demo

type User struct {
    Name string // 导出：别的包能读 u.Name
    age  int    // 未导出：只有本包能碰
}

func NewUser(name string) *User { // 导出的构造函数很常见
    return &User{Name: name, age: 0}
}
```

**对照 Java：** 没有 `public class` / `private` 关键字这套；也**没有**「子类可见的 protected」这条中间档。要分享给「少数内部包」用 `internal/`（见下），而不是 protected。

### 3. 目录 ≈ 边界，import 路径来自 module

- `go.mod` 里的 `module` 路径 + 目录 = import 路径。  
  例：`module github.com/you/app`，代码在 `internal/auth/` → `import "github.com/you/app/internal/auth"`。  
- **不要**把「一个业务功能拆成很多乱包」；包应是有凝聚度的 API 边界（后面 0501 再加深）。  
- 同一目录不能两个 package 名（测试文件 `package foo_test` 是刻意的外部测试包，以后 0309 再说）。

### 4. `internal/`：编译器强制的「私有子树」

- 路径里带 `internal` 的包：**只有** `internal` 的**父目录树之内**的代码能 import。  
- 例：`.../app/internal/auth` 可被 `.../app/...` 引用，**不能**被另一个 module 或无关兄弟工程引用。  

**对照 Java：** 有点像「禁止外人依赖你的内部实现包」，但是**工具链强制**，不是靠约定或 `module-info` 小心配置。

### 5. 和本课相关的小习惯

- 导出 API 保持小而稳；能小写就小写。  
- 需要给别的包用的构造/字段才大写。  
- 真正不想被别的 module 用的实现，放进 `internal/`。

### 6. 与后续 KP 的衔接

- **0102** 会在「类型与零值」上继续用导出类型。  
- **0401** 会把 `go.mod` / 依赖和 import 路径钉死。

## 最小例子

在脑子里过一遍（不必现在建 lab）：

```text
myapp/
  go.mod          // module example.com/myapp
  cmd/myapp/main.go    // package main；import example.com/myapp/user
  user/
    user.go            // package user；类型 User 导出，字段 id 不导出
  internal/store/
    store.go           // 只有 myapp 树内能 import
```

问自己：`other-module` 能不能 import `example.com/myapp/internal/store`？——不能。

## 验收题

用自己的话答（可打在聊天里）：

1. 为什么 `user.id` 小写、却常配一个导出的 `NewUser` / `ID()`？  
2. Java 的 `protected` 在 Go 里通常怎么替代（两种思路即可）？  
3. `internal` 包的限制是「约定」还是「编译器强制」？举一个谁能 import、谁不能的例子。

## 学完后 foundation 应如何改

- **已会** 增加：导出靠首字母；`internal/` 的强制可见性；目录与 import 路径关系。  
- **缺口** 去掉：「导出规则可能含糊」。  
- **仍开放：** 多包布局最佳实践留到 0501。

## 下一步

验收过关后：curriculum 0101 → `done`，更新 foundation，进入 **0102**。
