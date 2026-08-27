# Go 基础快照（Foundation）

> 写任何 `kp/*.md` **之前**读这份。  
> 学完一个 KP 并勾 `done` 时，**同一提交**更新本文件。  
> 知识目录：[curriculum.md](./curriculum.md)

- **最后更新：** 2026-08-27（试探初版）  
- **背景：** Java → Go；目标路径 C（先生产后端，再运行时深挖）  
- **当前焦点 KP：** 0106 slice header（建议；或先 P0 `java-to-go.md`）  
- **已完成 KP：** （无）

---

## 语言核心

**已会**

- 能读常见 Go 语法，能猜代码意图  
- 知道指针语法、slice/map 基本用法  
- 能写带 `go func` / `WaitGroup` 的结构  

**不稳 / 缺口**

- **逃逸分析：** 试探题 1 按 C 悬空指针理解；需建立「返回局部指针可安全」  
- **slice 共享底层数组：** 试探题 2 答错；header（ptr/len/cap）需专练 → **0106 优先**  
- **interface nil：** 试探题 3 不会；(type,value) 二元组未建立 → **0113 优先**  
- error 包装（Is/As）、接收者选择、嵌入组合：未测，默认待补  

**易混（Java 迁移）**

- 对象/引用 ≈ 到处堆上 → Go 常栈上，逃逸才堆  
- `List.subList` / 拷贝直觉 → slice 视图常共享数组  
- `null` → 带类型的 nil 仍可能让 interface ≠ nil  

---

## 并发

**已会**

- 知 goroutine；知闭包抓循环变量「都可能」不对劲（试探题 4）  
- context：级联取消 + 请求级数据（试探题 5，并类比 ThreadLocal）  

**不稳 / 缺口**

- Go 1.22+ 循环变量语义 vs 旧版本  
- `-race` 实操、channel 关闭约定、泄漏排查  
- 锁 vs channel 的选用、errgroup、常见模式落地  

**易混（Java 迁移）**

- `Thread` / 线程池成本 ≠ goroutine  
- `ThreadLocal` → context.Value 要克制、显式传递  

---

## 标准库

**已会**

- HTTP/context 方向正确；有后端工程直觉  

**不稳 / 缺口**

- 未系统覆盖：`io` 组合、httptest、slog、database/sql、embed 等（见 curriculum 03）  

---

## 工具链

**已会**

- 预期会 `go build` / `go test` 级基础（未细测）  

**不稳 / 缺口**

- vet/staticcheck/race、work、generate、模块 replace/版本 —— 默认待补  

---

## 生产工程化

**已会**

- Java 服务经验可迁移（分层、中间件、统一错误等概念）  

**不稳 / 缺口**

- Go 味布局、显式中间件链、优雅关机、依赖「先标准库」选型 —— 待 P3  

**易混（Java 迁移）**

- 默认上框架 / MDC / 全局异常处理 → 在 Go 里要改成显式、小接口、直接映射  

---

## 运行时与性能

**已会**

- （无专项）  

**不稳 / 缺口**

- GMP、GC、pprof、对着项目做逃逸/分配优化 —— 全部在 P4/06  

---

## 更新记录

| 日期 | 变更 |
|------|------|
| 2026-08-27 | 初版：五道试探题 + Java 背景写入 |
