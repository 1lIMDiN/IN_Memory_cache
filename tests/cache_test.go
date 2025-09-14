package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"cache/service"
)

func TestCacheConcurrency(t *testing.T) {
	cache := service.NewCache(time.Minute, time.Minute, 1000)
	var wg sync.WaitGroup

	// Тест на конкурентность
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			cache.Set(key, i, time.Second * 10)
			if _, found := cache.Get(key); !found {
				t.Errorf("Key %s not found", key)
			}
		}(i)
	}
	wg.Wait()
}