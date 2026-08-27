# Go 学习路线图（专属）

> 背景：有一定 Go 手感，主语言经历是 **Java**。目标路径 **C**：先能独立扛生产向后端，再啃语言/运行时。
>
> 原则：不从零语法课开刷；专治已测出的缺口；全程对照 Java（像什么 / 哪里不像）。

## 文档分工

| 文档 | 职责 |
|------|------|
| [curriculum.md](./curriculum.md) | **全量知识点 SSOT**（编号、前置/引出、状态、验收） |
| [foundation.md](./foundation.md) | **活基础快照**（讲解只认这份；学完必更新） |
| 本文件 roadmap | **阶段地图**（P0–P4 节奏、lab/project 挂载） |
| `kp/*.md` | 逐点讲解（**教完再写**；模板见 [_kp-template.md](./_kp-template.md)） |
| demo / lab | 实践时写可运行小文件，不空讲 |

进度以 curriculum 状态为准。活水位以 foundation 为准（下面表格是试探归档，不日常改）。

## 水位快照（试探归档）

| 区域 | 判断 | 对策 |
|------|------|------|
| 语法手感 | 能读能猜 | P0 快速对齐即可 |
| 指针 / 逃逸 | 偏 C/Java 对象直觉 | P1 → KP 0104 |
| slice | 底层数组共享不熟 | P1 → KP 0106 首攻 |
| interface / nil | 明显缺口 | P1 → KP 0113 |
| 并发 | 有直觉，细节/版本差待补 | P2 → 02 域 |
| HTTP / context | 工程直觉在线 | P3 → 03/05 域 |
| Java 迁移 | ThreadLocal / 线程池等可迁移 | 每 KP 对照 + 0509 |

## 阶段总览

```text
P0 心智对齐 ──► P1 语言补洞 ──► P2 并发打牢 ──► P3 生产后端 ──► P4 精通向深挖
   (短)            (优先)           (打底)           (主线干活)         (吃透)
```

**与 curriculum 域映射：** P0 ≈ 对照笔记 + 01 前段心智；P1 ≈ **01**；P2 ≈ **02**；P3 ≈ **03–05**；P4 ≈ **06**（可少量交叉）。

---

## P0 · Java → Go 心智对齐（短）

**目标：** 少用 Java 习惯硬套，建立 Go 默认心智。

| 主题 | Java 里像… | Go 里要记住 |
|------|------------|-------------|
| 包与可见性 | `public` / package-private | 大写导出、小写包内；目录 ≈ 包 |
| 错误 | 受检异常 / RuntimeException | 错误是值：`error`，多返回值 |
| 类型与复用 | class + 继承 | struct + 组合；没有继承树 |
| 并发单元 | `Thread` / 线程池 | `goroutine` 很轻；别按「一请求一线程」成本估算 |
| 接口 | `interface` 要显式 `implements` | 隐式满足；用法和 nil 坑与 Java 不同 |
| 请求级数据 | `ThreadLocal` | `context.Context`（更显式，别当万能袋子） |

**产出：** 一篇对照笔记即可（可放 `notes/java-to-go.md`）。  
**完成标准：** 能用自己的话讲清上表每一行「哪里不像」。

---

## P1 · 语言补洞（优先）

**目标：** 把试探里错的/不会的变成肌肉记忆。

1. **slice / map**
   - slice header（ptr/len/cap）与共享 backing array
   - `append` 是否会换底层数组；`copy` 的时机
   - map 不是线程安全；逗号 ok、delete、遍历顺序
2. **逃逸与指针**
   - 何时上堆；返回局部变量指针为什么在 Go 里安全
   - 和 Java「对象都在堆上」的差异
3. **interface**
   - `(type, value)` 二元组；**带类型的 nil ≠ nil interface**
   - 小接口、接受 interface 返回具体类型等惯用法
4. **error**
   - `errors.Is` / `errors.As` / `%w` 包装
   - 对照 Java：不是异常栈默认路径，但可保留因果链
5. **泛型**
   - 够用即可：约束、常见集合/工具函数；不深挖类型集哲学

**Lab：** `labs/001-slice-and-interface/`  
- 复现「改子切片影响原切片」与「nil interface」两个实验，写成 README 结论  

**完成标准：** 能独立讲清试探题 1～3 的正确答案与原因。

---

## P2 · 并发打牢

**目标：** 写出不易 race 的并发代码，并会用工具验证。

1. goroutine 生命周期与泄漏
2. channel / select / 关闭约定
3. `sync.WaitGroup`、`Mutex`、何时用 channel 何时用锁
4. `go test -race`；Go 1.22+ 循环变量语义变化
5. `context` 取消与超时；`errgroup`
6. 对照 Java：线程池、`CompletableFuture`、`ExecutorService.shutdown`

**Lab：** `labs/002-worker-pool/`  
- 带超时的 worker pool；故意造一个 race 再用 `-race` 抓住  

**完成标准：** 能解释试探题 4 在 1.21 vs 1.22 的差异，并给出版本无关的稳妥写法。

---

## P3 · 生产向后端（主线干活）

**目标：** 能独立交付一个可维护的小服务。

1. `net/http`：路由、middleware、优雅关机
2. 贯穿使用 `r.Context()`（取消 + 适度 request-scoped value）
3. 配置（env / 文件）、结构化日志
4. 测试：table-driven、handler 测试、`httptest` / 集成测试边界
5. modules、版本、基础依赖管理
6. 温和可观测：请求日志、健康检查；指标/追踪点到为止

**Project：** `projects/json-api/`  
- 小 JSON API：CRUD + 超时/取消 + 中间件 + 测试  

**完成标准：** 别人能按 README 跑起来、跑测试，并看懂目录职责。

---

## P4 · 精通向深挖

**目标：** 遇到性能/疑难时知道往哪看。

1. `pprof`（CPU / heap / goroutine）
2. GC 观感与逃逸实战（对着自己的 P3 项目画像）
3. 调度与 GMP 概览（概念层，不要求背源码）
4. 精读一块标准库：`net/http` 或 `sync`（二选一先深）

**Lab：** `labs/003-pprof-json-api/`  
- 对 P3 项目做一次性能画像，笔记里写「瓶颈假设 → 证据 → 改动」  

**完成标准：** 能对一个慢点给出有数据的解释，而不是只靠猜。

---

## 建议节奏

| 阶段 | 大致体量 | 备注 |
|------|----------|------|
| P0 | 0.5～1 天 | 写一篇对照笔记就过 |
| P1 | 数天～1 周 | 最优先；和试探缺口对齐 |
| P2 | 约 1 周 | 与 P1 可少量交错，但 race 要会 |
| P3 | 1～2 周 | 主线项目，可迭代加功能 |
| P4 | 持续 | 有真实问题再挖，避免空学运行时 |

日常习惯：笔记进 `notes/`，实验进 `labs/`，项目进 `projects/`；阶段状态可在本文件标题旁或主题 README 里改。

## 当前进度

阶段 checkbox（细项以 curriculum 为准）：

- [ ] P0
- [ ] P1
- [ ] P2
- [ ] P3
- [ ] P4

**下一步：** 按 curriculum **编号顺序** —— 当前 **0101**。
