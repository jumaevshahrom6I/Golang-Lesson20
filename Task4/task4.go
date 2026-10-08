// 4. Запроси https://jsonplaceholder.typicode.com/todos
// клиентом с таймаутом 5 секунд. Выведи, сколько задач
// выполнено (completed: true), а сколько нет.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(
		"https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	var todos []Todo
	err = json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	done := 0
	for _, t := range todos {
		if t.Completed {
			done++
		}
	}

	fmt.Println("Completed:", done)
	fmt.Println("Not completed:", len(todos)-done)
}
