// 15. Запроси https://jsonplaceholder.typicode.com/todos.
// Для каждого userId по возрастанию выведи количество
// выполненных и невыполненных задач, а в конце выведи
// id пользователей, у которых выполнено больше
// половины задач.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type stat struct {
	done    int
	notDone int
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

	// count per userId
	stats := map[int]*stat{}
	for _, t := range todos {
		s, ok := stats[t.UserID]
		if !ok {
			s = &stat{}
			stats[t.UserID] = s
		}
		if t.Completed {
			s.done++
		} else {
			s.notDone++
		}
	}

	// map has no order: collect keys and sort
	ids := make([]int, 0, len(stats))
	for id := range stats {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	var topUsers []int
	for _, id := range ids {
		s := stats[id]
		fmt.Printf("userId %d: done %d, not done %d\n",
			id, s.done, s.notDone)

		if s.done*2 > s.done+s.notDone {
			topUsers = append(topUsers, id)
		}
	}

	fmt.Println()
	fmt.Println("More than half done:", topUsers)
}
