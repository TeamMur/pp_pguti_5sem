package main

// 3.	Асинхронная обработка клиентских соединений:
//   •	Добавьте в сервер многопоточную обработку нескольких клиентских соединений.
//   •	Используйте горутины для обработки каждого нового соединения.
//   •	Реализуйте механизм graceful shutdown: сервер должен корректно завершать все активные соединения при остановке.

import (
	"fmt"
)

func task_3() {
	fmt.Println("hi")
}
