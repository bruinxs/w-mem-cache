package mcache

import (
	"time"

	"github.com/silenceper/wechat/cache"
)

//Default global cache
var Default cache.Cache = NewMemoryCache()

//Item value
type Item struct {
	Val      interface{}
	ExpireAt time.Time
}

//MemoryCache cache implement
type MemoryCache struct {
	m map[string]*Item
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		m: make(map[string]*Item, 3),
	}
}

func (mc *MemoryCache) Get(key string) interface{} {
	if !mc.IsExist(key) {
		return nil
	}

	return mc.m[key].Val
}

func (mc *MemoryCache) Set(key string, val interface{}, timeout time.Duration) error {
	mc.m[key] = &Item{
		Val:      val,
		ExpireAt: time.Now().Add(timeout),
	}
	return nil
}

func (mc *MemoryCache) IsExist(key string) bool {
	if item, ok := mc.m[key]; ok {
		if time.Now().Before(item.ExpireAt) {
			return true
		}
	}
	return false
}

func (mc *MemoryCache) Delete(key string) error {
	if item, ok := mc.m[key]; ok {
		item.ExpireAt = time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	}
	return nil
}
