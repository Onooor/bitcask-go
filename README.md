# Bitcask-Go: 轻量级嵌入式 KV 存储引擎

🚀 项目简介
Bitcask-Go 是基于 Bitcask 论文 设计的键值存储引擎，纯 Go 实现。支持 String/Hash/Set/List/SortedSet 五种 Redis 风格数据结构，结合内存索引的高效性与磁盘持久化的可靠性，适用于嵌入式存储场景。

✨ 核心特性

多数据结构支持

String/Hash/Set/List/SortedSet 完整实现示例：// Hash 操作
engine.HSet("user:1", map[string][]byte{"name": []byte("Alice")})


高效内存索引

抽象索引接口（BTree/ART/SkipList 可插拔）并发安全设计，支持高吞吐读写
持久化与崩溃恢复

数据追加写入日志文件自动恢复机制保障数据安全一键备份：engine.Backup("/backup_path")
HTTP 接口

RESTful API 远程操作：# 写入数据
curl -X PUT http://localhost:8080/bitcask/key -d 'value'


跨平台支持

兼容 Linux/macOS/Windows


⚡ 快速开始
安装
go get github.com/yourusername/bitcask-go

基本操作
package main

import (
	"github.com/yourusername/bitcask-go"
)

func main() {
	// 初始化引擎
	opts := bitcask.DefaultOptions
	opts.DirPath = "/tmp/bitcask-data"
	engine, _ := bitcask.Open(opts)
	defer engine.Close()

	// String 操作
	engine.Put([]byte("name"), []byte("Bitcask-Go"))
	val, _ := engine.Get([]byte("name")) // "Bitcask-Go"

	// Set 操作
	engine.SAdd([]byte("myset"), []byte("member1"))
	exists := engine.SIsMember([]byte("myset"), []byte("member1")) // true

	// 启动 HTTP 服务
	httpServer := bitcask.NewHTTPServer(engine, ":8080")
	httpServer.Start()
}

数据结构支持概览



类型
命令
示例




String
Put/Get/Delete
engine.Put(key, value)


Hash
HSet/HGet/HDel
engine.HSet(key, field, value)


Set
SAdd/SIsMember
engine.SAdd(key, member)


List
LPush/RPush/LPop
engine.LPush(key, value)


SortedSet
ZAdd/ZScore
engine.ZAdd(key, member, score)




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


📂 数据存储结构
/tmp/bitcask-data
├── data-0001.log    # 活跃数据文件
├── data-0002.log    # 归档文件
├── index.meta       # 索引元数据
└── lock             # 进程锁文件


📚 设计文档

核心设计

内存与磁盘设计数据备份机制
数据结构实现

String 结构支持Hash 结构支持Set 结构支持List 结构支持SortedSet 结构支持
扩展功能

HTTP 接口设计基准测试方案


🤝 参与贡献

提交 Issue 报告问题Fork 项目并提交 Pull Request遵循 Go 代码规范
# 运行测试
go test -v ./...


📜 许可证
MIT License © 2023 Your Name

项目地址
🔗 github.com/yourusername/bitcask-go


轻量如羽，持久如岩 —— 为你的应用提供可靠存储！

