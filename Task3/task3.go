// 3. Запроси по очереди
// https://jsonplaceholder.typicode.com/todos/5 и
// https://jsonplaceholder.typicode.com/todos/99999.
// Для каждого выведи статус-код. Если статус 200 —
// выведи title задачи, если 404 — выведи
// "Задача не найдена".

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/99999",
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Request error:", err)
			continue
		}

		fmt.Println("Status code:", resp.StatusCode)

		switch resp.StatusCode {
		case http.StatusOK:
			var todo Todo
			err := json.NewDecoder(resp.Body).Decode(&todo)
			if err != nil {
				fmt.Println("Parse error:", err)
			} else {
				fmt.Println("Title:", todo.Title)
			}
		case http.StatusNotFound:
			fmt.Println("Task not found")
		}

		resp.Body.Close()
		fmt.Println()
	}
}
