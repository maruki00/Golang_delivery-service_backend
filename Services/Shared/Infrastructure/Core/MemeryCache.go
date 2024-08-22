package shared_core

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	exprires time.Time
	value    interface{}
}

type MemoryCache struct {
	sync.Mutex
	items map[string]*CacheItem
}

func (o *MemoryCache) Insert(key string, expiresInSeconds int, val interface{}) error {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	o.items[key] = &CacheItem{
		value:    val,
		exprires: time.Now().Local().Add(time.Second * time.Duration(expiresInSeconds)),
	}
	return nil
}

func (o *MemoryCache) Delete(key string) {
	delete(o.items, key)
}

func (o *MemoryCache) Update(key string, expires int, val interface{}) error {
	_, ok := o.items[key]
	if ok {
		return fmt.Errorf("key not found")
	}

	if err := o.Insert(key, expires, val); err != nil {
		return err
	}
	return nil
}

func (o *MemoryCache) Clear() {
	o.items = make(map[string]*CacheItem)
}
