package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Set(key, val []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	keyStr := string(key)

	c.data[keyStr] = val

	return nil
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	keyStr := string(key)

	val, ok := c.data[keyStr]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", keyStr)
	}

	return val, nil
}

func (c *Cache) Exist(key []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.data[string(key)]

	return ok
}

func (c *Cache) Remove(key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	keyStr := string(key)

	delete(c.data, keyStr)

	return nil
}
