# Bitcask-Go: 轻量级嵌入式 KV 存储引擎

[![Go Report Card](https://goreportcard.com/badge/github.com/Onooor/bitcask-go)](https://goreportcard.com/report/github.com/Onooor/bitcask-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🚀 项目简介
**Bitcask-Go** 是基于 [Bitcask 论文](https://riak.com/assets/bitcask-intro.pdf) 设计的键值存储引擎，纯 Go 实现。支持 **String/Hash/Set/List/SortedSet** 五种 Redis 风格数据结构，结合内存索引的高效性与磁盘持久化的可靠性，适用于嵌入式存储场景。

---

## ✨ 核心特性

### 1. 多数据结构支持
- **String**: `Put`/`Get`/`Delete`
- **Hash**: `HSet`/`HGet`/`HDel`
- **Set**: `SAdd`/`SIsMember`/`SRem`
- **List**: `LPush`/`RPush`/`LPop`/`RPop`
- **SortedSet**: `ZAdd`/`ZScore`/`ZPopMax`

```go
// Hash 操作示例
engine.HSet("user:1", map[string][]byte{
    "name": []byte("Alice"),
    "age":  []byte("30")
})
