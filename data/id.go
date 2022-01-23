package data

import "time"

type Id struct {
	// 由什么创建而来，db/redis
	CreateType string
	// 类型 数字型、字符型
	Type string
	// 数字型id
	ValueInt int64
	// 字符型id
	ValueStr string
	// 创建时间
	CreateTime time.Time
}

type IdCreateParam struct {
	Node       NodeInfo
	CreateType string
	IdType     string
	// 步长
	Step int64
}

type NodeInfo struct {
	// Ip
	Ip string
	// 名称
	Name string
}
