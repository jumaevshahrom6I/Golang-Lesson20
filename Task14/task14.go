// 14. Для поста с id 1 выведи отчёт: заголовок поста,
// имя автора (пост хранит только userId, имя нужно
// запросить отдельно), количество комментариев к посту
// и email каждого комментатора. Все запросы выполняются
// последовательно одним клиентом с таймаутом 5 секунд.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const base = "https://jsonplaceholder.typicode.com"

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func getJSON(c *http.Client, url string, out any) error {
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	// 1) post
	var post Post
	err := getJSON(client, base+"/posts/1", &post)
	if err != nil {
		fmt.Println("Post error:", err)
		return
	}

	// 2) author, by userId from the post
	var user User
	userURL := fmt.Sprintf("%s/users/%d", base, post.UserID)
	err = getJSON(client, userURL, &user)
	if err != nil {
		fmt.Println("Author error:", err)
		return
	}

	// 3) post comments
	var comments []Comment
	err = getJSON(client, base+"/posts/1/comments", &comments)
	if err != nil {
		fmt.Println("Comments error:", err)
		return
	}

	fmt.Println("Title:", post.Title)
	fmt.Println("Author:", user.Name)
	fmt.Println("Comments:", len(comments))
	for _, c := range comments {
		fmt.Println(" -", c.Email)
	}
}
