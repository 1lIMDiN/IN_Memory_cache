package main

import (
	"fmt"
	"time"

	"cache/service"
)

func main() {
	// Создаем кэш с TTL 7 минут, очисткой в 1 минуту, размер 125
	cache := service.NewCache(7 * time.Minute, 1 * time.Minute, 125)

	// Добавляем значения
	cache.Set("DNS:admin", "admin`-`", 25 * time.Second)
	cache.Set("DNS:user", 25, 0)

	// Получаем значения 
	if value, found := cache.Get("DNS:user"); found {
		fmt.Println("User:", value)
	}

	// Обновляем значения
	cache.Set("DNS:user", "Ivan Ivanov", 33 * time.Second)

	// Получаем значения 
	if value, found := cache.Get("DNS:user"); found {
		fmt.Println("User:", value)
	}

	// Проверяем существование ключа
	if cache.Exists("DNS:admin") {
		fmt.Println("Admin exists")
	}

}