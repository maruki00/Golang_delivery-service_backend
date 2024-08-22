package shared_core

import (
	"sync"
	"time"
)

type CacheItem struct {
	exprires time.Time
	value    any
}

type MemoryCache struct {
	sync.Mutex
	items map[string]CacheItem
}

func (o *MemoryCache) Insert(key string, val any) error {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()

	return nil
}

func (o *MemoryCache) Delete() error {

}

func (o *MemoryCache) Update() error {

}

func (o *MemoryCache) Insert() error {

}
