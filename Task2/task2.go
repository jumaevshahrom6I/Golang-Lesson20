// 2. Отправь GET-запрос на
// https://jsonplaceholder.typicode.com/users/1
// клиентом с таймаутом 5 секунд. Разбери ответ в
// структуру с полями id, name, email, phone и выведи
// имя, email и телефон.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(
		"https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Phone:", user.Phone)
}
