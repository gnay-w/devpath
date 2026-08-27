# Conventions

## 布局

- 顶层按技术/主题分目录：`topics/<tech>/`
- 每个主题**永远**有三槽：`notes/`、`labs/`、`projects/`（空也保留）
- 根目录只放导航与约定；真实内容进 `topics/`

## 命名

- 技术目录与 lab/project 子目录：小写 kebab-case（如 `system-design`、`001-hello-world`）
- `labs/` / `projects/` 建议用可选编号前缀，按加入顺序（不强制）
- 笔记文件名：日期或主题均可（如 `2026-08-27-ownership.md` / `ownership.md`），同一主题内自洽即可

## 三槽

| 槽 | 放什么 | 不放什么 |
|----|--------|----------|
| `notes/` | 概念笔记、读书摘录、对比总结、cheatsheet | 需要跑起来才能验证的东西 |
| `labs/` | 最小可运行实验；一问题一子目录；建议自带短 README（目的 + 怎么跑） | 多日迭代、有产品味道的练习 |
| `projects/` | 稍完整的练习/小项目 | 一次性试 API 的 hello-world |

边界：约 30 分钟内跑通 → `labs/`；要持续打磨 → `projects/`；只写不跑 → `notes/`。

## 主题 README

每个 `topics/<tech>/README.md` 包含：

1. **这是什么** — 一句话
2. **我为什么学** — 动机 / 目标
3. **状态** — `inbox` / `learning` / `parked` / `done`
4. **入口** — 重要笔记、代表性 lab、主项目（没有写「暂无」）
5. **下一步** — 当前 1～3 件事

根 `README.md` 技术表镜像：`技术名 | 状态 | 一句话`。

## 新建主题

```bash
cp -r _templates/topic topics/<tech>
# 编辑 topics/<tech>/README.md
# 在根 README.md 技术表加一行
```
