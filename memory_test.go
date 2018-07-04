package mcache

import (
	"testing"
	"time"
)

func TestMemoryCache(t *testing.T) {
	cache := NewMemoryCache()

	//get
	val := cache.Get("key")
	if val != nil {
		t.Error("val illegal")
		return
	}

	//set
	cache.Set("key", "name", time.Second*3)
	val = cache.Get("key")
	if s := val.(string); s != "name" {
		t.Error("error ", val)
		return
	}

	//expire
	time.Sleep(time.Second * 3)
	val = cache.Get("key")
	if val != nil {
		t.Error("val illegal")
		return
	}

	//delete
	cache.Set("key", "name", time.Minute*3)
	err := cache.Delete("key")
	if err != nil {
		t.Error(err)
		return
	}
	val = cache.Get("key")
	if val != nil {
		t.Error("illegal ", val)
		return
	}
}
