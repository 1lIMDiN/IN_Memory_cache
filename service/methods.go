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
		CleanupInterval:   cleanupInterval,
		MaxSize:           maxSize,
		Stop:              make(chan bool),
	}

	if cleanupInterval > 0 {

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

// Keys возвращает все ключи
func (c *Cache) Keys() []string {
	c.RLock()
	defer c.RUnlock()

	keys := make([]string, 0, len(c.Entries))
	now := time.Now().UnixNano()

	for k, v := range c.Entries {
		if v.Expiration == 0 || now <= v.Expiration {
			keys = append(keys, k)
		}
	}

	return keys
}

// Сount возвращает кол-во актульных элементов
func (c *Cache) Count() int {
	c.RLock()
	defer c.RUnlock()

	count := 0
	now := time.Now().UnixNano()

	for _, v := range c.Entries {
		if v.Expiration == 0 || now <= v.Expiration {
			count++
		}
	}

	return count
}
// deleteOld удаляет старые данные
func (c *Cache) deleteOld() {

}
