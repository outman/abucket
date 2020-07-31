package pkg

/*
Copyright © 2020 pochonlee@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

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
