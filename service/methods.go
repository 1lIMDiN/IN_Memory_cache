package service

import (
	"errors"
	"time"
)

// NewCache создает новый кэш
func NewCache(defaultExpiration, cleanupInterval time.Duration, maxSize int) *Cache {
	cache := &Cache{
		Entries:           make(map[string]Entry),
		DefaultExpiration: defaultExpiration,
		CleanuoInterval:   cleanupInterval,
		MaxSize:           maxSize,
		Stop:              make(chan bool),
	}

	return cache
}

// Set добавляет/обновляет значение
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	// Освобождаем место, если достигли лимита и ключ новый
	if c.MaxSize > 0 && len(c.Entries) >= c.MaxSize && c.Entries[key] == (Entry{}) {
		c.deleteOld()
	}

	var expiration int64
	if ttl == 0 {
		ttl = c.DefaultExpiration
	}
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}

	c.Entries[key] = Entry{
		Value:      value,
		Created:    time.Now(),
		Expiration: expiration,
	}
}

// Get возвращает значение по ключу
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.Entries[key]
	if !found {
		return nil, false
	}

	// Проверка TTL
	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

// Delete удаляет значение по ключу
func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.Entries[key]; !found {
		return errors.New("there is nothing to delete")
	}

	delete(c.Entries, key)
	return nil
}

// Exists проверяeт существование ключа
func (c *Cache) Exists(key string) bool {
	c.RLock()
	defer c.RUnlock()

	item, found := c.Entries[key]
	if !found {
		return false
	}

	// Проверка TTL
	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		return false
	}

	return true
}

// deleteOld удаляет старые данные
func (c *Cache) deleteOld() {

}
