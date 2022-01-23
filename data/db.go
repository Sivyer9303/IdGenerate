package data

import (
	"sync"
	"time"
)

var (
	// 单机锁
	mu = sync.Mutex{}
)

func init() {

}

type IdInfo struct {
	Id         int64 `gorm:"primaryKey"`
	BizType    string
	BeginId    int64
	maxId      int64
	Step       int64
	Delta      int64
	Remainder  int64
	CreateTime time.Time
	UpdateTime time.Time
	Version    string
}
