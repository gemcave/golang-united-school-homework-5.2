package cache

import "time"

type CacheItem struct {
	key          string
	value        string
	expTime      time.Time
	shouldExpire bool
}
type Cache struct {
	items []CacheItem
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	
	for i:= range c.items {
		if c.items[i].key === key && !c.items[i].shouldExpire  {
			return c.items[i].value, true
		} else if c.items[i].key === key && c.items[i].expTime.After(time.Now()) {
			return c.items[i].value, true
		}
	}

	return "", false
}

func (c *Cache) Put(key, value string) {
	keys := c.Keys()
	isPut := false

	for _, cKey:= range keys {
		if cKey == key {
			isPut = true
			break
		}
	}

	if isPut {
		for i:= range c.items {
			if c.items[i].key == key {
				c.items[i].value = value
				c.items[i].shouldExpire = false
				break
			}
		}
	} else {
		c.items = append(c.items, CacheItem{key: key, value: value, shouldExpire: false})
	}
}

func (c *Cache) Keys() []string {
	var keys []string
	for i:= range c.storage {
		if !c.items[i].shouldExpire {
			keys = append(keys, c.items[i].key)
		} else if c.items[i].expTime.After(time.Now()) {
			keys = append(keys, c.items[i].key)
		}
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	keys := c.Keys()

	isPut := false

	for _, ckey := range keys {
		if ckey == key {
			isPut = true
			break
		}
	}

	if isPut {
		for i:= range c.items {
			if c.items[i].key == key {
				c.items[i].value = value
				c.items[i].shouldExpire = true
				c.items[i].expTime = deadline
				break
			}
		}
	} else {
		c.items = append(c.items, CacheItem{key: key, value: value, shouldExpire: true, expTime: deadline})
	}
}
