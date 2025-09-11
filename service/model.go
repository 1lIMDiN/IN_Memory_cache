package service

import (
	"sync"
	"time"
)

// Cashe - кэш с TTL
type Cashe struct {
	sync.RWMutex
	DefaultExpiration time.Duration
	CleanuoInterval   time.Duration
	MaxSize           int
	Entries           map[string]Entry
	Stop              chan bool
}

// Entry - запись в кэше
type Entry struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

