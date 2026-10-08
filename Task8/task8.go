// 8. Запроси https://httpbin.org/delay/2 клиентом с
// таймаутом 1 секунда — выведи полученную ошибку.
// Затем отправь тот же запрос клиентом с таймаутом
// 5 секунд — выведи статус-код.

package main

import (
	"fmt"
	"net/http"
	"time"
)

const apiURL = "https://httpbin.org/delay/2"

func request(timeout time.Duration) {
	client := &http.Client{Timeout: timeout}

	resp, err := client.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
}

func main() {
	fmt.Println("Timeout 1 second:")
	request(1 * time.Second)

	fmt.Println()

	fmt.Println("Timeout 5 seconds:")
	request(5 * time.Second)
}
