package pkg

import (
	"sync"
	"time"
)

type cacheItem struct {
	object     interface{}
	expiration int64
}

type cacheStore struct {
	item     *sync.Map
	duration int64
}

// NewCacheStore (storage *sync.Map, duration int64)
// type cacheStore struct {
// 	Item     *sync.Map
// 	Duration int64
// }
func NewCacheStore(storage *sync.Map, duration int64) *cacheStore {
	return &cacheStore{
		item:     storage,
		duration: duration,
	}
}

func (c *cacheStore) Get(k interface{}) (interface{}, bool) {
	v, ok := c.item.Load(k)
	if ok {
		i := v.(cacheItem)
		if i.expiration == 0 || i.expiration > time.Now().Unix() {
			return i.object, ok
		}
	}
	c.item.Delete(k)
	return nil, false
}

func (c *cacheStore) Set(k interface{}, v interface{}) {
	item := cacheItem{
		object:     v,
		expiration: c.duration + time.Now().Unix(),
	}
	c.item.Store(k, item)
}

func (c *cacheStore) Delete(k interface{}) {
	c.item.Delete(k)
}
