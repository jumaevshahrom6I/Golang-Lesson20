// 10. Отправь POST-запрос на
// https://jsonplaceholder.typicode.com/posts клиентом с
// таймаутом 5 секунд. Тело — JSON с полями title, body и
// userId (равен 1). Выведи статус-код и id созданного
// поста из ответа.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NewPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type Created struct {
	ID int `json:"id"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	post := NewPost{
		Title:  "My first post",
		Body:   "Post text",
		UserID: 1,
	}

	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Encode error:", err)
		return
	}

	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	var created Created
	err = json.NewDecoder(resp.Body).Decode(&created)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Created post ID:", created.ID)
}
