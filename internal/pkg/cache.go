package pkg

import (
	"sync"
	"time"
)

type cacheItem struct {
	object     interface{}
	expiration int64
}

type CacheStore struct {
	Item     sync.Map
	Duration int64
}

func (c *CacheStore) Get(k interface{}) (interface{}, bool) {
	v, ok := c.Item.Load(k)
	if ok {
		i := v.(cacheItem)
		if i.expiration == 0 || i.expiration > time.Now().Unix() {
			return i.object, ok
		}
	}
	c.Item.Delete(k)
	return nil, false
}

func (c *CacheStore) Set(k interface{}, v interface{}) {
	item := cacheItem{
		object:     v,
		expiration: c.Duration + time.Now().Unix(),
	}
	c.Item.Store(k, item)
}

func (c *CacheStore) Delete(k interface{}) {
	c.Item.Delete(k)
}
