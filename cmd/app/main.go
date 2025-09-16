package main

import (
	"fmt"
	"time"

	"cache/service"
)

func main() {
	// Создаем кэш с TTL 7 минут, очисткой в 1 минуту, размер 125
	cache := service.NewCache(7*time.Minute, 1*time.Minute, 125)

	// Добавляем значения
	cache.Set("DNS:admin", "admin`-`", 5*time.Second)
	cache.Set("DNS:user", 25, 0)

	// Получаем значения для User
	if value, found := cache.Get("DNS:user"); found {
		fmt.Println("User:", value)
	}

	// Обновляем значения для User
	cache.Set("DNS:user", "Ivan Ivanov", 33*time.Second)

	// Получаем новое значения для User
	if value, found := cache.Get("DNS:user"); found {
		fmt.Println("User:", value)
	}

	// Проверяем существование ключа
	if cache.Exists("DNS:admin") {
		fmt.Println("Admin exists")
	}

	// Получаем все ключи
	keys := cache.Keys()
	for i := 0; i < len(keys); i++ {
		fmt.Printf("Key %d: %v\n", i+1, keys[i])
	}

	// Проверка TTL
	time.Sleep(1 * time.Second) // Установи 5 или более секунд
	if v, ok := cache.Get("DNS:admin"); !ok {
		fmt.Println("DNS:admin expired")
	} else {
		fmt.Println("DNS:admin exists:", v)
	}

	// Удаляем Admin и проверяем количество актуальных элементов
	fmt.Println("Count before delete:", cache.Count())
	cache.Delete("DNS:admin")
	fmt.Println("Count after delete:", cache.Count())

	// Останавливаем чистку
	defer cache.StopClean()
	// Чистим весь кэш
	defer cache.FlushAll()
}
