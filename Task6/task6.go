// 6. Получи посты пользователя с userId=3 (адрес
// https://jsonplaceholder.typicode.com/posts, userId
// передай как query-параметр). Выведи количество постов
// и title каждого.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

	params := url.Values{}
	params.Set("userId", "3")

	apiURL := "https://jsonplaceholder.typicode.com/posts" +
		"?" + params.Encode()

	resp, err := client.Get(apiURL)
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

	fmt.Println("Posts count:", len(posts))
	for i, p := range posts {
		fmt.Printf("%d. %s\n", i+1, p.Title)
	}
}
