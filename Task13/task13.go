// 13. Напиши структуру APIClient с полем baseURL и
// http-клиентом с таймаутом 5 секунд, конструктор
// NewAPIClient(baseURL string) и два метода:
// GetUser(id int) и GetUserPosts(userID int). Если
// сервер вернул статус, отличный от 200, метод
// возвращает ошибку с этим статусом. В main создай
// клиент с baseURL https://jsonplaceholder.typicode.com,
// получи пользователя 1 и его посты, выведи имя и
// количество постов. Затем запроси пользователя 11 и
// выведи полученную ошибку.

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
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
}

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

// get is a helper: request, status check,
// and JSON decoding
func (c *APIClient) get(path string, out any) error {
	resp, err := c.client.Get(c.baseURL + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *APIClient) GetUser(id int) (*User, error) {
	var user User
	err := c.get(fmt.Sprintf("/users/%d", id), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *APIClient) GetUserPosts(id int) ([]Post, error) {
	var posts []Post
	path := fmt.Sprintf("/posts?userId=%d", id)
	err := c.get(path, &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func main() {
	api := NewAPIClient("https://jsonplaceholder.typicode.com")

	user, err := api.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	posts, err := api.GetUserPosts(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Posts count:", len(posts))
	fmt.Println()

	_, err = api.GetUser(11)
	if err != nil {
		fmt.Println("User 11 error:", err)
	}
}
