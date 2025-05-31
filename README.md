# bitcask-go
Bitcask-Go: 轻量级嵌入式 KV 存储引擎

🚀 项目简介
Bitcask-Go 是一个基于 Bitcask 论文 设计的轻量级键值存储引擎，使用纯 Go 实现。它结合了内存索引的高效性和磁盘持久化的可靠性，支持多种数据结构（String/Hash/Set/List/ZSet），并提供 HTTP 接口、数据备份等实用功能。
架构示意图
架构图：内存索引 + 磁盘数据文件

✨ 核心特性

多数据结构支持

String/Hash/Set/List/SortedSet 五种 Redis 风格数据结构。示例：engine.HSet("user:1", map[string][]byte{"name": []byte("Alice"), "age": []byte("30")})


高效内存索引

抽象索引接口（BTree/ART/SkipList 可插拔）。并发安全，支持高吞吐读写。
持久化与备份

数据追加写入日志文件，崩溃后自动恢复。一键备份：engine.Backup("/backup_path")。
HTTP 接口

通过 RESTful API 远程操作：curl -X PUT http://localhost:8080/bitcask/key -d 'value'


跨平台

支持 Linux/macOS/Windows。


⚡ 快速开始
安装
go get github.com/yourusername/bitcask-go

基本用法
package main

import (
	"github.com/yourusername/bitcask-go"
)

func main() {
	// 打开存储引擎
	opts := bitcask.DefaultOptions
	opts.DirPath = "/tmp/bitcask-data"
	engine, _ := bitcask.Open(opts)
	defer engine.Close()

	// 写入数据
	engine.Put([]byte("name"), []byte("Bitcask-Go"))

	// 读取数据
	val, _ := engine.Get([]byte("name"))
	fmt.Println(string(val)) // 输出: Bitcask-Go

	// 使用 Hash 结构
	engine.HSet("user:1", map[string][]byte{"email": []byte("user@example.com")})
}

启动 HTTP 服务
httpServer := bitcask.NewHTTPServer(engine, ":8080")
httpServer.Start()


📊 性能基准测试



操作
吞吐量 (ops/sec)
平均延迟 (μs)




Put
78,000
12.8


Get
120,000
8.3


Delete
95,000
10.5




测试环境：4vCPU/8GB RAM/SSD 磁盘


📂 数据目录结构
/tmp/bitcask-data
├── data-0001.log    # 活跃数据文件
├── data-0002.log    # 归档数据文件
├── index.meta       # 索引元数据
└── lock             # 文件锁


📚 文档导航

设计细节

内存与磁盘设计Redis 数据结构支持
扩展功能

HTTP 接口数据备份机制


🤝 参与贡献

提交 Issue 报告问题或建议。Fork 项目并提交 Pull Request。遵循 Go 代码规范。
# 运行测试
go test -v ./...


📜 许可证
MIT License © 2023 Your Name

项目地址
🔗 github.com/yourusername/bitcask-go


轻量如羽，持久如岩 —— 为你的应用提供可靠存储！


