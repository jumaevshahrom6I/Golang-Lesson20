// 16. Запроси пользователей с id от 1 до 10
// (https://jsonplaceholder.typicode.com/users/)
// одновременно через горутины и WaitGroup, каждая
// горутина кладёт результат в свою ячейку среза. Выведи
// имена по порядку id и общее время работы. Затем
// выполни те же 10 запросов последовательно и выведи
// время работы обоих вариантов рядом.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const base = "https://jsonplaceholder.typicode.com/users/"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func fetchName(c *http.Client, id int) string {
	resp, err := c.Get(fmt.Sprintf("%s%d", base, id))
	if err != nil {
		return "error: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("error: status %d", resp.StatusCode)
	}

	var u User
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return "error: " + err.Error()
	}
	return u.Name
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	const n = 10

	// --- parallel ---
	names := make([]string, n)
	var wg sync.WaitGroup

	start := time.Now()
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// each goroutine writes only to its own cell,
			// so no mutex is needed
			names[id-1] = fetchName(client, id)
		}(i)
	}
	wg.Wait()
	parallel := time.Since(start)

	fmt.Println("Parallel:")
	for i, name := range names {
		fmt.Printf("%d. %s\n", i+1, name)
	}
	fmt.Println("Time:", parallel)
	fmt.Println()

	// --- sequential ---
	start = time.Now()
	for i := 1; i <= n; i++ {
		fetchName(client, i)
	}
	sequential := time.Since(start)

	fmt.Printf("Parallel:   %v\n", parallel)
	fmt.Printf("Sequential: %v\n", sequential)
}
