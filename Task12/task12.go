// 12. Запроси https://jsonplaceholder.typicode.com/users
// клиентом с таймаутом 5 секунд. Для каждого
// пользователя выведи name и название его компании
// (company.name).

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Company struct {
	Name string `json:"name"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Company Company `json:"company"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(
		"https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	for _, u := range users {
		fmt.Printf("%s - %s\n", u.Name, u.Company.Name)
	}
}
