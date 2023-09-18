# go-structs-kit

go-structs-kit 是一个用Go语言编写的数据结构SDK，主要包括优先队列以及预留空间用于未来添加其他高效的数据结构。

## 特点

- 基于Go1.18泛型实现
- 并发安全
- 高效、可扩展
- 简单易用的API


## 安装

使用Go的包管理工具进行安装：

```bash
go get github.com/jaylee630/go-structs-kit
```

## 使用方法

### 优先队列
```go
import "github.com/jaylee630/go-structs-kit/pkg/priority_queue"

// 创建优先队列
pq := priority_queue.New[int]()

// 添加元素
pq.Push(Item[int]{Value: 1, Priority: 2})

// 弹出元素
item := pq.Pop()
```

更多使用示例，请查看 examples/ 目录。

## 文档

详细文档请查看 docs/ 目录。

## 贡献

欢迎提交Pull Requests或提出Issues。

## 许可证

本项目采用MIT许可证，详见 `LICENSE` 文件。
