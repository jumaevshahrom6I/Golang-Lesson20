// 5. Запроси https://jsonplaceholder.typicode.com/posts
// клиентом с таймаутом 5 секунд. Выведи количество
// постов, а также title первого и последнего поста.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(
		"https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	if len(posts) == 0 {
		fmt.Println("No posts")
		return
	}

	fmt.Println("Posts count:", len(posts))
	fmt.Println("First:", posts[0].Title)
	fmt.Println("Last:", posts[len(posts)-1].Title)
}
