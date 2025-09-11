package main

import (
	"fmt"
	"time"

	"cashe/service"
)

func main() {
	// Создаем кэш с TTL 7 минут, очисткой в 1 минуту, размер 125
	cashe := service.NewCashe(7 * time.Minute, 1 * time.Minute, 125)

	// Добавляем значения
	cashe.Set("DNS:admin", "admin`-`", 25 * time.Second)
	cashe.Set("DNS:user", 25, 0)

	// Получаем значения 
	if value, found := cashe.Get("DNS:user"); found {
		fmt.Println("User:", value)
	}

	// Обновляем значения
	cashe.Set("DNS:user", "Ivan Ivanov", 33 * time.Second)

	// Получаем значения 
	if value, found := cashe.Get("DNS:user"); found {
		fmt.Println("User:", value)
	}

	// Проверяем существование ключа
	if cashe.Exists("DNS:admin") {
		fmt.Println("Admin exists")
	}

}